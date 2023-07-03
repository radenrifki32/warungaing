package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rifki321/warungku/product"
	"github.com/rifki321/warungku/user"
)

type Order struct {
	Id               int
	NoPemesanan      string
	TanggalPemesanan time.Time
	ProductId        int
	TotalBarang      int
	Product          product.ProductEntity
	User             user.User
}
type OrderImpl interface {
	OrderProduct(ctx context.Context, tx *sql.Tx, ProductId int, totalBarang int, userId int) (Order, error)
}

type OrderRepoImpl struct {
}

func NewRepoOrder() *OrderRepoImpl {
	return &OrderRepoImpl{}
}

func (orderRepo *OrderRepoImpl) OrderProduct(ctx context.Context, tx *sql.Tx, ProductId int, totalBarang int, userId int) (Order, error) {
	OrderNoPemesanan := uuid.Must(uuid.NewRandom())
	fmt.Println(ProductId)
	fmt.Println(totalBarang)
	fmt.Println(userId)
	sqlStatementForInsert := `
	INSERT INTO orders (no_pemesanan, product_id, total_barang, user_id)
	SELECT tanggal_pemesanan,no_pemesanan,
	WHERE EXISTS (SELECT 1 FROM products WHERE id = ? AND quantity_product > 0)
`
	sqlStatementForUpdate := `
UPDATE products SET quantity_product = quantity_product - ? WHERE id = ? 
`
	sqlStatementForQuery := "SELECT p.nama_product, p.harga_product, p.quantity_product, c.tanggal_pemesanan, c.no_pemesanan, u.username FROM products AS p JOIN orders AS c ON p.id = c.product_id JOIN users AS u ON c.user_id = u.id WHERE c.order_id = ?"
	result, err := tx.ExecContext(ctx, sqlStatementForInsert, OrderNoPemesanan.String(), ProductId, totalBarang, userId, ProductId)
	if err != nil {
		return Order{}, err
	}
	OrderId, err := result.LastInsertId()

	if err != nil {
		return Order{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return Order{}, err
	}

	if rowsAffected == 0 {
		return Order{}, err
	}
	_, errUpdate := tx.ExecContext(ctx, sqlStatementForUpdate, totalBarang, ProductId)
	if errUpdate != nil {
		return Order{}, err
	}

	fmt.Println(rowsAffected)
	rows, err := tx.QueryContext(ctx, sqlStatementForQuery, int(OrderId))
	if err != nil {
		return Order{}, err
	}
	defer rows.Close()
	var order Order
	for rows.Next() {
		var NamaProduct string
		var HargaProduct string
		var Quantity int
		var TanggalPemesanan time.Time
		var NoPemesanan string

		err := rows.Scan(&NamaProduct, &HargaProduct, &Quantity, &TanggalPemesanan, &NoPemesanan)
		if err != nil {
			return Order{}, err
		}

		order.Product = product.ProductEntity{
			NamaProduct:  NamaProduct,
			HargaProduct: HargaProduct,
			Quantity:     int32(Quantity),
		}
		order.TanggalPemesanan = TanggalPemesanan
		order.NoPemesanan = NoPemesanan
	}
	return order, nil

}
