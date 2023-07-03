package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepo interface {
	Register(ctx context.Context, sql *sql.Tx, user User) (User, error)
	Login(ctx context.Context, sql *sql.Tx, username string) (User, error)
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
		return User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	user.Id = int32(id)
	return user, nil

}

func (userRepo *UserImplementation) Login(ctx context.Context, sql *sql.Tx, username string) (User, error) {
	sqlQuery := "SELECT username,password FROM users WHERE username = ? LIMIT 1"
	rows, err := sql.QueryContext(ctx, sqlQuery, username)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()
	userlogin := User{}
	if rows.Next() {
		if err := rows.Scan(&userlogin.Username, &userlogin.Password); err != nil {
			return User{}, err
		}
		fmt.Println(userlogin.Password)
		return userlogin, nil
	} else {
		return User{}, errors.New("User Not Found")
	}
}
