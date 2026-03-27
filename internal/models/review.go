package models

type Reviews struct {
	FullName string  `json:"full_name"`
	Messages string  `json:"messages"`
	Rating   int     `json:"rating"`
	Picture  *string `json:"picture,omitzero"`
}
