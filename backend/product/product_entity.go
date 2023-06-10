package product

import "github.com/rifki321/warungku/categories/entity"

type ProductEntity struct {
	Id           int32
	IdCategory   int
	NamaProduct  string
	HargaProduct string
	Quantity     int32
	Category     entity.Category
}
