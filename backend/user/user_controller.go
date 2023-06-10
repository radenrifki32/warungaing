package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/user/web"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	service UserService
}

func NewUserController(service UserService) UserController {
	return &UserControllerImpl{service: service}
}
func (controller *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RegisterRequest := &web.RegisterRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(RegisterRequest)
	if err != nil {
		panic(err)
	}

	UserResponse := controller.service.Register(r.Context(), *RegisterRequest)

	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "Ok",
		Data:   UserResponse,
	}
	fmt.Println(webResponse)
	WriteFromJsonBody(w, webResponse)

}
func (controller *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RegisterRequest := &web.RegisterRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(RegisterRequest)
	if err != nil {
		fmt.Println("Decode Not Valid")
	}
	fmt.Println(&RegisterRequest.Username)
	helper.LogBody(w, r)

	UserResponse := controller.service.Login(r.Context(), *RegisterRequest)

	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "Ok",
		Data:   UserResponse,
	}
	WriteFromJsonBody(w, webResponse)

}

func WriteFromJsonBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)

	if err != nil {
		panic(err)

	}
}
