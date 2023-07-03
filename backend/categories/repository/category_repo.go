package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rifki321/warungku/categories/entity"
)

type CategoryInterface interface {
	GetCategory(ctx context.Context, tx *sql.Tx) ([]entity.Category, error)
}
type CategoryStruct struct {
}

func NewCategory() *CategoryStruct {
	return &CategoryStruct{}
}

func (repo *CategoryStruct) GetCategory(ctx context.Context, tx *sql.Tx) ([]entity.Category, error) {
	query := "SELECT id,nama_category from category"
	rows, err := tx.QueryContext(ctx, query)
	if err := rows.Close(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		fmt.Println(err)

		return nil, err
	}
	var categories []entity.Category
	for rows.Next() {
		categori := entity.Category{}
		if err := rows.Scan(&categori.Id, &categori.NamaCategory); err != nil {
			fmt.Println(err)

			return nil, err
		}
		categories = append(categories, categori)

	}
	return categories, nil

}
