package web

import "github.com/rifki321/warungku/categories/entity"

type ResponseProduct struct {
	Id           int32
	NamaProduct  string
	HargaProduct string
	Quantity     int32
}

type ResponseProductWithCategory struct {
	Id           int32
	NamaProduct  string
	HargaProduct string
	Quantity     int32
	Category     entity.Category
}
