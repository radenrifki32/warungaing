package response

import (
	"github.com/rifki321/warungku/categories/entity"
)

type Response struct {
	NamaCategory string `json:"nama_category"`
}

type Request struct {
	Id int32
}

type ResponseCategoryWeb struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseCategory(category []entity.Category) []Response {
	var categories []Response
	for _, res := range category {
		categories = append(categories, SingleResponseCategori(res))
	}
	return categories
}

func SingleResponseCategori(category entity.Category) Response {
	return Response{NamaCategory: category.NamaCategory}
}
