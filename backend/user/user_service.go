package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rifki321/warungku/app"
	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/user/web"
)

type UserService interface {
	Register(ctx context.Context, request web.RegisterRequest) web.ResponseLogin
	Login(ctx context.Context, request web.RegisterRequest) web.ResponseLogin
}

type UserServiceImpl struct {
	repo UserRepo
	sql  *sql.DB
}

func NewUserService(repo UserRepo, sql *sql.DB) UserService {
	return &UserServiceImpl{repo: repo, sql: sql}
}
func (service *UserServiceImpl) Register(ctx context.Context, request web.RegisterRequest) web.ResponseLogin {

	fmt.Println(request)
	tx, err := service.sql.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	hashingPassword, err := app.HashPassword(request.Password)
	if err != nil {
		panic(err)
	}
	user := User{
		Username: request.Username,
		Password: hashingPassword,
	}
	user, err = service.repo.Register(ctx, tx, user)
	if err != nil {
		panic(web.ResponseErrorNotFound(err.Error()))
	}
	return ToResponseUser(user, "")
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.RegisterRequest) web.ResponseLogin {
	tx, err := service.sql.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitOrRollback(tx)
	user := User{
		Username: request.Username,
		Password: request.Password,
	}
	user, err = service.repo.Login(ctx, tx, user)
	if err != nil {
		panic(web.ResponseErrorNotFound(err.Error()))
	}
	token, err := app.GenerateToken(user.Username)
	if err != nil {
		fmt.Println(err)
	}

	err = app.OpenPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(web.ResponseErrorNotFound(err.Error()))
	}

	return ToResponseUser(user, token)

}

func ToResponseUser(user User, token string) web.ResponseLogin {
	return web.ResponseLogin{
		Username: user.Username,
		Token:    token,
	}

}
