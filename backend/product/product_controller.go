package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/product/webproduct"
	"github.com/rifki321/warungku/user/web"
)

type ProductController interface {
	GetAllProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetProductById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	PostProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetProductByCategoriesId(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetProductByCategories(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type ProductControllerImpl struct {
	service ProductService
}

func NewProductController(service ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{service: service}
}

func (service *ProductControllerImpl) GetAllProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ResponseProduct := service.service.GetAllProduct(r.Context(), w)
	helper.Response(w, http.StatusOK, "OK", ResponseProduct)

}
func (controller *ProductControllerImpl) GetProductById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getId := params.ByName("productid")
	id, err := strconv.Atoi(getId)
	if err != nil {
		panic(err)
	}
	ResponseProduct := controller.service.GetProductById(r.Context(), w, int32(id))
	webResponse := web.ResponseWeb{
		Code:   http.StatusOK,
		Status: true,
		Data:   ResponseProduct,
	}
	WriteFromJsonBody(w, webResponse)

}

func (controller *ProductControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decode := json.NewDecoder(r.Body)
	ProductRequest := webproduct.RequestProduct{}
	decode.Decode(&ProductRequest)
	product := controller.service.PostProduct(r.Context(), w, ProductRequest)
	responseWeb := web.ResponseProduct{
		NamaProduct:  product.NamaProduct,
		HargaProduct: product.HargaProduct,
		Quantity:     product.Quantity,
	}
	WriteFromJsonBody(w, responseWeb)

}

func (controller *ProductControllerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("productid")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	controller.service.DeleteProduct(r.Context(), int32(idProduct))

}

func (controller *ProductControllerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decode := json.NewDecoder(r.Body)
	ProductRequest := webproduct.RequestProduct{}
	err := decode.Decode(&ProductRequest)
	if err != nil {
		fmt.Println("gagal decode")
	}
	id := params.ByName("productid")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	ProductRequest.Id = int32(idProduct)
	fmt.Println(ProductRequest.Id)
	product := controller.service.UpdateProduct(r.Context(), w, ProductRequest)

	webResponse := web.ResponseWeb{
		Code:   200,
		Status: true,
		Data:   product,
	}
	fmt.Println(product)
	WriteFromJsonBody(w, webResponse)
}

func (controller *ProductControllerImpl) GetProductByCategoriesId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("categoryid")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	responseProduct := controller.service.GetProductByCategoriesId(r.Context(), categoryId)
	responseWeb := web.ResponseWeb{
		Code:   200,
		Status: true,
		Data:   responseProduct,
	}
	WriteFromJsonBody(w, responseWeb)
}
func (controller *ProductControllerImpl) GetProductByCategories(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	responseProduct := controller.service.GetProductByCategories(r.Context())
	responseWeb := web.ResponseWeb{
		Code:   200,
		Status: true,
		Data:   responseProduct,
	}
	WriteFromJsonBody(w, responseWeb)
}

func WriteFromJsonBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)

	if err != nil {
		panic(err)

	}
}
