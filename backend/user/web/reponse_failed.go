package web

import (
	"encoding/json"
	"net/http"
)

func ResponseFailed(w http.ResponseWriter, request *http.Request, err interface{}) {
	notFoundError(w, request, err)
	internalServerError(w, request, err)
}
func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	responseWeb := ResponseWeb{
		Code:   http.StatusInternalServerError,
		Status: false,
		Data:   err,
	}
	WriteFromJsonBody(w, responseWeb)

}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(ErrorNotFounded)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		responseWeb := ResponseWeb{
			Code:   http.StatusNotFound,
			Status: false,
			Data:   exception.Error,
		}
		WriteFromJsonBody(w, responseWeb)
		return true
	} else {
		return false
	}
}
func WriteFromJsonBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)

	if err != nil {
		panic(err)

	}
}
