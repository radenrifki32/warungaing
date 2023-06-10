package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepo interface {
	Register(ctx context.Context, sql *sql.Tx, user User) (User, error)
	Login(ctx context.Context, sql *sql.Tx, user User) (User, error)
}
type UserImplementation struct {
}

func NewUserRepoRepository() UserRepo {
	return &UserImplementation{}
}

func (userRepo *UserImplementation) Register(ctx context.Context, sql *sql.Tx, user User) (User, error) {
	SqlQuery := "insert into users(username,password) values(?,?)"
	result, err := sql.ExecContext(ctx, SqlQuery, user.Username, user.Password)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.Id = int32(id)
	fmt.Println("selesai")
	return user, nil

}

func (userRepo *UserImplementation) Login(ctx context.Context, sql *sql.Tx, user User) (User, error) {
	sqlQuery := "SELECT username, password FROM users WHERE username = ?"
	rows, err := sql.QueryContext(ctx, sqlQuery, user.Username)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	userlogin := User{}
	if rows.Next() {
		rows.Scan(&userlogin.Username, &userlogin.Password)
		fmt.Println(userlogin)
		return userlogin, nil
	} else {
		return userlogin, errors.New("User Not Found")
	}
}
