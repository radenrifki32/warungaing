package user

import "time"

type User struct {
	Id        int32
	Username  string
	Password  string
	CreatedAt time.Time
}
