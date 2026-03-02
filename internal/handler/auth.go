package handler

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/arif14377/koda-b6-backend/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var listUsers []entity.Users
var conn *pgx.Conn

// start - connection database
func InitDB() {
	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		fmt.Println("Failed to parse config")
		return
	}

	conDb, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		fmt.Println("Failed to connect to db")
		return
	}

	conn = conDb
	// end - connection database

}

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

	data.Id = len(listUsers) + 1
	listUsers = append(listUsers, data)
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
			if u.Email == data.Email && u.Password == data.Password {
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
	notFound := true

	for i, u := range listUsers {
		if u.Id == id {
			listUsers = slices.Delete(listUsers, i, i+1)
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
