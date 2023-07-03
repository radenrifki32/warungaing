package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rifki321/warungku/categories/repository"
	"github.com/rifki321/warungku/categories/response"
	"github.com/rifki321/warungku/helper"
)

type ServiceCategory interface {
	GetCategory(ctx context.Context) ([]response.Response, error)
}

type ServiceCategoryStruct struct {
	repo repository.CategoryInterface
	sql  *sql.DB
}

func NewServiceCategory(repo repository.CategoryInterface, sql *sql.DB) *ServiceCategoryStruct {
	return &ServiceCategoryStruct{repo: repo, sql: sql}
}

func (service *ServiceCategoryStruct) GetCategory(ctx context.Context) ([]response.Response, error) {
	tx, err := service.sql.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)
	category, err := service.repo.GetCategory(ctx, tx)
	fmt.Println(category)
	if err != nil {
		return nil, err
	}
	return response.ResponseCategory(category), nil
}
