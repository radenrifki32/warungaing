package entitymessage

import (
	"time"

	"github.com/rifki321/warungku/user"
)

type Message struct {
	Id          int       `json:"id"`
	Subject     string    `json:"subject_message"`
	Fromto      string    `json:"user_from"`
	SenderName  string    `json:"sender_from"`
	Message     string    `json:"message"`
	Datemessage time.Time `json:"tanggal_message"`
	User        user.User `json:"user"`
	TotalRow    int       `json:"total_row"`
}
