package controllerOrder

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/app/order/serviceOrder"
	"github.com/rifki321/warungku/helper"
)

type OrderController interface {
	OrderProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type OrderControllerImpl struct {
	service serviceOrder.OrderService
}

func NewControllerOrder(service serviceOrder.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{service: service}
}

func (controller *OrderControllerImpl) OrderProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	OrderRequest := serviceOrder.OrderRequest{}
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&OrderRequest)
	if err != nil {
		helper.Response(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR", nil)
		return
	}
	response, err := controller.service.OrderProduct(r.Context(), OrderRequest)
	if err != nil {
		helper.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	fmt.Println(response)
	helper.Response(w, http.StatusOK, "Success", response)
}

func WriteFromJsonBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)

	if err != nil {
		helper.Response(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR", nil)
	}
}
