package user

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/user/web"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
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
	registerRequest := web.RegisterRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&registerRequest)
	if err != nil {
		helper.Response(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR", nil)
		return
	}

	userResponse, err := controller.service.Register(r.Context(), registerRequest)
	if err != nil {
		helper.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusCreated, "Success", userResponse)

}
func (controller *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loginRequest := web.LoginRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginRequest)
	if err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if loginRequest.Username == "" || loginRequest.Password == "" {
		helper.Response(w, http.StatusBadRequest, "Username atau Password tidak boleh kosong", nil)
		return
	}

	userResponse, err := controller.service.Login(r.Context(), loginRequest, w)
	if err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Success", userResponse)
}

func WriteFromJsonBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)

	if err != nil {
		panic(err)

	}
}
