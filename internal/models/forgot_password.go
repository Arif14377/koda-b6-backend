package models

import "math/big"

type ForgotPassword struct {
	Id              int      `json:"id"`
	Email           string   `json:"email"`
	Code            *big.Int `json:"code"`
	NewPassword     string   `json:"newPassword"`
	ConfirmPassword string   `json:"confirmPassword"`
}

type VerifOTP struct {
	Email string   `json:"email"`
	Code  *big.Int `json:"code"`
}
