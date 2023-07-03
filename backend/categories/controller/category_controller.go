package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/categories/response"
	"github.com/rifki321/warungku/categories/service"
)

type ContollerContract interface {
	GetCategory(ctx context.Context)
}

type ControllerCategory struct {
	service service.ServiceCategory
}

func NewControllerCategory(service service.ServiceCategory) *ControllerCategory {
	return &ControllerCategory{service: service}
}
func (controller *ControllerCategory) GetCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp, err := controller.service.GetCategory(r.Context())
	fmt.Println(resp)
	if err != nil {
		webResponse := response.ResponseCategoryWeb{
			Status: "Failed",
			Code:   http.StatusInternalServerError,
		}
		WriteFromJsonBody(w, webResponse, http.StatusInternalServerError)
		return
	}
	webResponse := response.ResponseCategoryWeb{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   resp,
	}
	WriteFromJsonBody(w, webResponse, http.StatusOK)

}

func WriteFromJsonBody(w http.ResponseWriter, response interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)

	if err != nil {
		panic(err)

	}
}
