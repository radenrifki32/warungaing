package serviceOrder

import (
	"context"
	"database/sql"
	"time"

	"github.com/rifki321/warungku/app/order/repo"
	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/product"
	"github.com/rifki321/warungku/user"
)

type OrderRequest struct {
	ProductId   int `json:"product_id"`
	TotalBarang int `json:"total_barang"`
	UserId      int `json:"user_id"`
}
type OrderResponse struct {
	NoPemesanan      string                `json:"no_pemesanan"`
	TanggalPemesanan time.Time             `json:"tanggal_pemesanan"`
	Product          product.ProductEntity `json:"product"`
	User             user.User             `json:"user"`
}

type OrderService interface {
	OrderProduct(ctx context.Context, requstOrder OrderRequest) (responseOrder OrderResponse, err error)
}
type OrderServiceImpl struct {
	sql       *sql.DB
	OrderRepo repo.OrderImpl
}

func NewOrderService(sql *sql.DB, OrderRepo repo.OrderImpl) *OrderServiceImpl {
	return &OrderServiceImpl{sql: sql, OrderRepo: OrderRepo}
}
func (service *OrderServiceImpl) OrderProduct(ctx context.Context, requestOrder OrderRequest) (responseOrder OrderResponse, err error) {
	tx, err := service.sql.Begin()
	if err != nil {
		return OrderResponse{}, err

	}
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepo.OrderProduct(ctx, tx, requestOrder.ProductId, requestOrder.TotalBarang, requestOrder.UserId)
	if err != nil {
		return OrderResponse{}, err
	}
	responseWebOrder := OrderResponse{
		NoPemesanan:      order.NoPemesanan,
		TanggalPemesanan: order.TanggalPemesanan,
		Product:          order.Product,
		User:             order.User,
	}
	return responseWebOrder, err
}
