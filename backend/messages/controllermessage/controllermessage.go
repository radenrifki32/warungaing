package controllermessage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/app"
	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/messages/servicemessage"
)

type ControllerMessage interface {
	SendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetMessageByUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetMessageByIdMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
type ControllerMessageImpl struct {
	SeviceMessage servicemessage.ServiceMessage
}

func NewControllerMessage(ServiceMessage servicemessage.ServiceMessage) *ControllerMessageImpl {
	return &ControllerMessageImpl{SeviceMessage: ServiceMessage}
}
func (serviceMessage *ControllerMessageImpl) SendMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jwtCookie := r.Header.Get("Cookie")
	fmt.Println(jwtCookie)
	jwtSplit := strings.Split(jwtCookie, "=")
	jwtString := jwtSplit[1]
	if jwtString == "" {
		helper.Response(w, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}
	token, err := app.VerifyToken(strings.TrimPrefix(jwtString, "Bearer "))
	if err != nil {
		helper.Response(w, http.StatusUnauthorized, "UNAUTHORIZATION", nil)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		helper.Response(w, http.StatusUnauthorized, "UNAUTHORIZATION", nil)
		return
	}
	username := claims["username"].(string)
	RequestMessage := servicemessage.RequestMessage{}
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&RequestMessage); err != nil {
		helper.Response(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR", nil)
		return
	}
	response, err := serviceMessage.SeviceMessage.SendMessage(r.Context(), &RequestMessage, username)
	if err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helper.Response(w, http.StatusOK, "OK", response)

}
func (serviceMessage *ControllerMessageImpl) GetMessageByUsername(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jwtCookie := r.Header.Get("Cookie")
	jwtSplit := strings.Split(jwtCookie, "=")
	jwtString := jwtSplit[1]
	if jwtString == "" {
		helper.Response(w, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}
	token, err := app.VerifyToken(strings.TrimPrefix(jwtString, "Bearer "))
	if err != nil {
		helper.Response(w, http.StatusUnauthorized, "UNAUTHORIZATION", nil)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {

		helper.Response(w, http.StatusUnauthorized, "UNAUTHORIZATION", nil)
		return
	}
	username := claims["username"].(string)
	responseMessage, totalRow, err := serviceMessage.SeviceMessage.GetMessageByUsername(r.Context(), username)
	fmt.Println(totalRow)
	if err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response := struct {
		Messages  []servicemessage.ResponseMessageByUsername `json:"messages"`
		TotalRows int                                        `json:"total_rows"`
	}{
		Messages:  responseMessage,
		TotalRows: totalRow,
	}

	helper.Response(w, http.StatusOK, "OK", response)

}

func (messageController *ControllerMessageImpl) GetMessageByIdMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("messageId")
	params, err := strconv.Atoi(id)
	if err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	message, err := messageController.SeviceMessage.GetMessageByIdMessage(r.Context(), params)
	if err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helper.Response(w, http.StatusOK, "BERHASIL", message)

}
