package entity

import "time"

type Users struct {
	Id         int       `json:"id"`
	FullName   string    `json:"fullName"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	Photo      string    `json:"photo"`
	Role       string    `json:"role"`
	Created_At time.Time `json:"created_at"`
	Created_By *string   `json:"created_by"`
	Updated_At time.Time `json:"updated_at"`
	Updated_By *string   `json:"updated_by"`
}

type UserListRead struct {
	Id       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}
