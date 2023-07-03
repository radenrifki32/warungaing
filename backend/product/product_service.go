package product

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/product/webproduct"
	"github.com/rifki321/warungku/user/web"
)

type ProductService interface {
	GetAllProduct(ctx context.Context, w http.ResponseWriter) []web.ResponseProduct
	GetProductById(ctx context.Context, w http.ResponseWriter, ProductId int32) web.ResponseProduct
	PostProduct(ctx context.Context, w http.ResponseWriter, r webproduct.RequestProduct) web.ResponseProduct
	DeleteProduct(ctx context.Context, productId int32)
	UpdateProduct(ctx context.Context, w http.ResponseWriter, r webproduct.RequestProduct) web.ResponseProduct
	GetProductByCategoriesId(ctx context.Context, categoryId int) []web.ResponseProductWithCategory
	GetProductByCategories(ctx context.Context) []web.ResponseProductWithCategory
}

type ProductServiceImpl struct {
	repo ProductsRepo
	db   *sql.DB
}

func NewProductService(repo ProductsRepo, db *sql.DB) *ProductServiceImpl {
	return &ProductServiceImpl{repo: repo, db: db}
}
func (service *ProductServiceImpl) GetAllProduct(ctx context.Context, w http.ResponseWriter) []web.ResponseProduct {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitOrRollback(tx)

	products, err := service.repo.GetAllProduct(ctx, tx)
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
	}
	return ToResponseProducts(products)

}

func (service *ProductServiceImpl) GetProductById(ctx context.Context, w http.ResponseWriter, ProductId int32) web.ResponseProduct {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	product, err := service.repo.GetProductById(ctx, tx, ProductId)
	if err != nil {
		panic(err)
	}
	return ToResponsProduct(product)

}
func (service *ProductServiceImpl) PostProduct(ctx context.Context, w http.ResponseWriter, r webproduct.RequestProduct) web.ResponseProduct {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	product := ProductEntity{
		Id:           r.Id,
		NamaProduct:  r.NamaProduct,
		HargaProduct: r.HargaProduct,
		Quantity:     r.Quantity,
	}
	service.repo.PostProduct(ctx, tx, product)
	return ToResponsProduct(product)
}

func (service *ProductServiceImpl) DeleteProduct(ctx context.Context, productId int32) {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	product, err := service.repo.GetProductById(ctx, tx, productId)
	if err != nil {
		panic(err)
	}
	service.repo.DeleteProduct(ctx, tx, product.Id)
}

func (service *ProductServiceImpl) UpdateProduct(ctx context.Context, w http.ResponseWriter, r webproduct.RequestProduct) web.ResponseProduct {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	product, err := service.repo.GetProductById(ctx, tx, r.Id)
	if err != nil {
	}
	product = ProductEntity{
		Id:           r.Id,
		NamaProduct:  r.NamaProduct,
		HargaProduct: r.HargaProduct,
		Quantity:     r.Quantity,
	}
	products := service.repo.UpdateProduct(ctx, tx, product)
	return ToResponsProduct(products)

}

func (service *ProductServiceImpl) GetProductByCategories(ctx context.Context) []web.ResponseProductWithCategory {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	product := service.repo.GetProductByCategories(ctx, tx)
	if err != nil {
		panic(err)
	}
	return ToResponseProductsWithCategory(product)

}
func (service *ProductServiceImpl) GetProductByCategoriesId(ctx context.Context, categoryId int) []web.ResponseProductWithCategory {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	product, err := service.repo.GetProductByCategoriesId(ctx, tx, categoryId)
	if err != nil {
		panic(err)
	}
	return ToResponseProductsWithCategory(product)

}
func ToResponseProducts(products []ProductEntity) []web.ResponseProduct {
	var responseProduct []web.ResponseProduct
	for _, product := range products {
		responseProduct = append(responseProduct, ToResponsProduct(product))
	}
	return responseProduct
}
func ToResponseProductsWithCategory(products []ProductEntity) []web.ResponseProductWithCategory {
	var responseProduct []web.ResponseProductWithCategory
	for _, product := range products {
		responseProduct = append(responseProduct, ToResponsProductWithCategory(product))
	}
	return responseProduct
}
func ToResponsProductWithCategory(product ProductEntity) web.ResponseProductWithCategory {
	return web.ResponseProductWithCategory{
		Id:           product.Id,
		NamaProduct:  product.NamaProduct,
		HargaProduct: product.HargaProduct,
		Quantity:     product.Quantity,
		Category:     product.Category,
	}

}
func ToResponsProduct(product ProductEntity) web.ResponseProduct {
	return web.ResponseProduct{
		Id:           product.Id,
		NamaProduct:  product.NamaProduct,
		HargaProduct: product.HargaProduct,
		Quantity:     product.Quantity,
	}

}
