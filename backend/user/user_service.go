package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rifki321/warungku/app"
	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/user/web"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, request web.RegisterRequest) (web.ResponseRegister, error)
	Login(ctx context.Context, request web.LoginRequest, w http.ResponseWriter) (web.ResponseLogin, error)
}

type UserServiceImpl struct {
	repo UserRepo
	sql  *sql.DB
}
type ErrorResponseMessage struct {
	Error string `json:"error,omitempty"`
	Code  int    `json:"code"`
}

func NewUserService(repo UserRepo, sql *sql.DB) UserService {
	return &UserServiceImpl{repo: repo, sql: sql}
}
func (serviceUser UserServiceImpl) Register(ctx context.Context, request web.RegisterRequest) (web.ResponseRegister, error) {
	tx, err := serviceUser.sql.Begin()
	if err != nil {
		return web.ResponseRegister{}, err
	}
	defer helper.CommitOrRollback(tx)
	if len(request.Password) < 8 {
		return web.ResponseRegister{}, errors.New("password kurang dari 8 karakter")
	}
	hashing, err := app.HashPassword(request.Password)
	user := User{
		Username:  request.Username,
		Password:  hashing,
		CreatedAt: time.Now(),
	}
	if err != nil {
		return web.ResponseRegister{}, err
	}
	User, err := serviceUser.repo.Register(ctx, tx, user)
	if err != nil {
		return web.ResponseRegister{}, err
	}
	response := web.ResponseRegister{
		Username:  User.Username,
		CreatedAt: time.Now(),
		Message:   "Register Success",
	}
	return response, nil
}

func (userService UserServiceImpl) Login(ctx context.Context, request web.LoginRequest, w http.ResponseWriter) (web.ResponseLogin, error) {
	tx, err := userService.sql.Begin()
	if err != nil {
		return web.ResponseLogin{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := userService.repo.Login(ctx, tx, request.Username)
	if err != nil {
		return web.ResponseLogin{}, err
	}
	fmt.Println(user, "user")

	if user.Username == "" {
		return web.ResponseLogin{}, errors.New("User Tidak Ditemukan")
	}
	if err := app.OpenPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return web.ResponseLogin{}, errors.New("username dan password salah")
		} else {
			return web.ResponseLogin{}, errors.New("Password tidak valid")
		}
	}

	token, err := app.GenerateToken(user.Username, w)
	if err != nil {
		return web.ResponseLogin{}, errors.New("error disini")
	}
	app.HashPassword(token)
	cookie := &http.Cookie{
		Name:     "jwt-token",
		Value:    token,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)

	response := web.ResponseLogin{
		Username:  user.Username,
		Token:     token,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}
