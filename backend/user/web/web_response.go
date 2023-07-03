package web

import "time"

type ResponseLogin struct {
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
}

type ResponseRegister struct {
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
