package models

import "time"

type Users struct {
	Id         string    `json:"id"`
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
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type UserEmail struct {
	Email string `json:"email"`
}
