package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rifki321/warungku/categories/entity"
)

type ProductsRepo interface {
	GetAllProduct(ctx context.Context, sql *sql.Tx) ([]ProductEntity, error)
	GetProductById(ctx context.Context, sql *sql.Tx, ProductId int32) (ProductEntity, error)
	PostProduct(ctx context.Context, sql *sql.Tx, product ProductEntity) (ProductEntity, error)
	DeleteProduct(ctx context.Context, sql *sql.Tx, productId int32)
	UpdateProduct(ctx context.Context, sql *sql.Tx, product ProductEntity) ProductEntity
	GetProductByCategoriesId(ctx context.Context, sql *sql.Tx, categoriesId int) ([]ProductEntity, error)
	GetProductByCategories(ctx context.Context, sql *sql.Tx) []ProductEntity
}
type ProductsRepoImpl struct {
}

func NewProductRepo() *ProductsRepoImpl {
	return &ProductsRepoImpl{}
}

func (prod *ProductsRepoImpl) GetAllProduct(ctx context.Context, sql *sql.Tx) ([]ProductEntity, error) {
	QueryProduct := "select nama_product,harga_product,quantity_product from products"
	rows, err := sql.QueryContext(ctx, QueryProduct)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var Products []ProductEntity
	for rows.Next() {
		product := ProductEntity{}
		err := rows.Scan(&product.NamaProduct, &product.HargaProduct, &product.Quantity)
		if err != nil {
			return nil, err

		}
		Products = append(Products, product)
	}
	return Products, nil

}

func (prod *ProductsRepoImpl) GetProductById(ctx context.Context, sql *sql.Tx, ProductId int32) (ProductEntity, error) {
	fmt.Println("repo", ProductId)
	QueryProductId := "SELECT id,nama_product,harga_product,quantity_product FROM products WHERE id = ?"
	rows, err := sql.QueryContext(ctx, QueryProductId, ProductId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	product := ProductEntity{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.NamaProduct, &product.HargaProduct, &product.Quantity)
		if err != nil {
			panic(err)
		}
		return product, nil

	}
	fmt.Println(product)
	return product, errors.New("category is not found")
}

func (prod *ProductsRepoImpl) PostProduct(ctx context.Context, sql *sql.Tx, product ProductEntity) (ProductEntity, error) {
	sqlQuery := "INSERT INTO products (nama_product,harga_product,quantity_product) VALUES(?,?,?)"
	result, err := sql.ExecContext(ctx, sqlQuery, product.NamaProduct, product.HargaProduct, product.Quantity)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	product.Id = int32(id)
	return product, nil

}

func (prod *ProductsRepoImpl) DeleteProduct(ctx context.Context, sql *sql.Tx, productId int32) {
	sqlQuery := "DELETE FROM products WHERE id = ?"
	_, err := sql.ExecContext(ctx, sqlQuery, productId)
	if err != nil {
		panic(err)
	}

}

func (prod *ProductsRepoImpl) UpdateProduct(ctx context.Context, sql *sql.Tx, product ProductEntity) ProductEntity {
	sqlQuery := "UPDATE products SET nama_product=?,harga_product=?,quantity_product =? WHERE id =?"
	_, err := sql.ExecContext(ctx, sqlQuery, product.NamaProduct, product.HargaProduct, product.Quantity, product.Id)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("berhasil")
	return product
}

func (prod *ProductsRepoImpl) GetProductByCategories(ctx context.Context, sql *sql.Tx) []ProductEntity {
	sqlQuery := "SELECT p.nama_product, p.harga_product, p.quantity_product, c.nama_category, c.id FROM products AS p JOIN category AS c ON (p.id_category = c.id)"
	rows, err := sql.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []ProductEntity
	for rows.Next() {
		product := ProductEntity{}
		category := entity.Category{}
		err := rows.Scan(&product.NamaProduct, &product.HargaProduct, &product.Quantity, &category.NamaCategory, &category.Id)
		if err != nil {
			panic(err)
		}
		product.Category = category
		products = append(products, product)
	}

	fmt.Println("slice", products)
	return products
}

func (prod *ProductsRepoImpl) GetProductByCategoriesId(ctx context.Context, sql *sql.Tx, categoriesId int) ([]ProductEntity, error) {
	fmt.Println("repo", categoriesId)
	sqlQuery := "SELECT p.nama_product, p.harga_product, p.quantity_product, c.nama_category, c.id FROM products AS p JOIN category AS c ON (p.id_category = c.id) WHERE c.id = ?"
	rows, err := sql.QueryContext(ctx, sqlQuery, categoriesId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []ProductEntity
	for rows.Next() {
		product := ProductEntity{}
		category := entity.Category{}
		err := rows.Scan(&product.NamaProduct, &product.HargaProduct, &product.Quantity, &category.NamaCategory, &category.Id)
		if err != nil {
			return nil, err
		}
		product.Category = category
		products = append(products, product)
	}

	fmt.Println("slice", products)
	return products, nil
}
