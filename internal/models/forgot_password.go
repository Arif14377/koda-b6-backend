package models

type ForgotPassword struct {
	Id              int      `json:"id"`
	Email           string   `json:"email"`
	Code            int      `json:"code"`
	NewPassword     string   `json:"newPassword"`
	ConfirmPassword string   `json:"confirmPassword"`
}

type VerifOTP struct {
	Email string   `json:"email"`
	Code  int      `json:"code"`
}
