package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/arif14377/koda-b6-backend/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/matthewhartstonge/argon2"
)

var listUsers []entity.Users
var conn *pgx.Conn
var argon argon2.Config

// start - connection database
func InitDB() error {
	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		return fmt.Errorf("Failed to parse config: %w\n", err)
	}

	conDb, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		return fmt.Errorf("Failed to connect to db: %w\n", err)
	}

	conn = conDb
	return nil
}

// end - connection database

// start - create JWT (Hmac)
func ClaimJWT(userID int) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"expired": time.Now().Add(time.Second * 20),
	})

	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))

	fmt.Println(tokenString, err)
}

// end - create JWT (Hmac)

func Register(ctx *gin.Context) {
	data := entity.Users{}
	err := ctx.ShouldBindJSON(&data)
	isExist := false

	if err != nil {
		ctx.JSON(401, entity.Response{
			Success: false,
			Message: "JSON tidak valid.",
		})
		return
	}

	// validasi email:
	// 1. email harus ada @
	// 2. fullname, email dan password tidak boleh kosong
	// 3. Jika email sudah terdaftar, maka tidak bisa register
	// 4. Selain itu register berhasil.

	if !strings.Contains(data.Email, "@") {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "Email tidak valid.",
		})
		return
	}

	if data.FullName == "" || data.Email == "" || data.Password == "" {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "Data tidak boleh kosong.",
		})
		return
	}

	for _, u := range listUsers {
		if data.Email == u.Email {
			isExist = true
		}
	}

	if isExist {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "Email sudah terdaftar.",
		})
		return
	}

	argon = argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(data.Password))
	if err != nil {
		log.Fatalf("Error encode %v\n", err)
		return
	}

	if conn == nil {
		ctx.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: "conn == nil",
		})
		return
	}

	data.Password = string(encoded)
	// data.Id = len(listUsers) + 1 //sudah increment dari DB

	// entity.Users
	// conn.Exec(ctx, `
	// CREATE TABLE IF NOT EXISTS (
	// 	id,
	// )
	// `)

	_, err = conn.Exec(context.Background(),
		`INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3)`, data.FullName, data.Email, data.Password,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// listUsers = append(listUsers, data) // upgraded
	ctx.JSON(200, entity.Response{
		Success: true,
		Message: "Registrasi berhasil.",
	})
}

func Login(ctx *gin.Context) {
	var data entity.Users
	err := ctx.ShouldBindJSON(&data)
	login := false

	if err != nil {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "JSON tidak valid",
		})
		return
	}

	if !strings.Contains(data.Email, "@") {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "Email tidak valid.",
		})
		return
	}

	if data.Email == "" || data.Password == "" {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "Email dan Password tidak boleh kosong.",
		})
		return
	}

	for _, u := range listUsers {
		if u.Email == data.Email {
			pwdOk, _ := argon2.VerifyEncoded([]byte(data.Password), []byte(u.Password))
			if u.Email == data.Email && pwdOk {
				login = true
			} else {
				ctx.JSON(400, entity.Response{
					Success: false,
					Message: "Password salah.",
				})
				return
			}
		}
	}

	if login {
		ctx.JSON(200, entity.Response{
			Success: true,
			Message: fmt.Sprintf("Welcome %s", data.Email),
		})
		ClaimJWT(data.Id)
	} else {
		ctx.JSON(401, entity.Response{
			Success: false,
			Message: "Email tidak terdaftar. Silahkan register terlebih dahulu.",
		})
	}
}

func GetUsers(ctx *gin.Context) {
	if conn == nil {
		ctx.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: "conn == nil",
		})
		return
	}

	rows, err := conn.Query(context.Background(),
		`SELECT id, full_name, email FROM users`,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: "gagal select users",
		})
		return
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[entity.UserListRead])

	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: "Failed to get data users",
		})
		return
	}

	ctx.JSON(200, entity.Response{
		Success: true,
		Message: "List Users:",
		Results: users,
	})

}

func UserDetails(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := entity.Users{}
	notFound := true

	for _, u := range listUsers {
		if u.Id == id {
			user = u
			notFound = false
			break
		}
	}

	if notFound {
		ctx.JSON(404, entity.Response{
			Success: false,
			Message: "User tidak ditemukan",
		})
		return
	}

	ctx.JSON(200, entity.Response{
		Success: true,
		Message: fmt.Sprintf("data user ID: %d", id),
		Results: user,
	})
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	// notFound := true

	// for i, u := range listUsers {
	// 	if u.Id == id {
	// 		listUsers = slices.Delete(listUsers, i, i+1)
	// 		notFound = false
	// 		break
	// 	}
	// }

	// if notFound {
	// 	ctx.JSON(404, entity.Response{
	// 		Success: false,
	// 		Message: "User tidak ditemukan",
	// 	})
	// 	return
	// }

	_, err := conn.Exec(context.Background(),
		`DELETE FROM users where id = $1`, id,
	)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}

	ctx.JSON(200, entity.Response{
		Success: true,
		Message: fmt.Sprintf("Data dengan id %d berhasil dihapus", id),
	})
}

func UpdateUser(ctx *gin.Context) {
	data := entity.Users{}
	err := ctx.ShouldBindJSON(&data)
	// Validasi update data:
	// 1. Email yang sudah terdaftar tidak bisa dipakai

	if err != nil {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "JSON tidak valid.",
		})
	}

	for _, u := range listUsers {
		if data.Email == u.Email {
			ctx.JSON(400, entity.Response{
				Success: false,
				Message: "Email sudah terdaftar.",
			})
			return
		}
	}

	listUsers = append(listUsers, data)
	ctx.JSON(200, entity.Response{
		Success: true,
		Message: "Data berhasil diperbarui.",
	})
}
