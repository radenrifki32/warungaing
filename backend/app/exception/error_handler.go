package exception

import (
	"encoding/json"
	"net/http"

	"github.com/rifki321/warungku/user/web"
)

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, error interface{}) {
	if ErrorNotFound(w, r, error) {
		return
	}
	internalServerError(w, r, error)
}

func ErrorNotFound(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		responseWeb := web.ResponseWeb{
			Status: true,
			Code:   http.StatusBadRequest,
			Data:   exception.Error,
		}
		WriteFromJsonBody(w, responseWeb)
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.ResponseWeb{
		Code:   http.StatusInternalServerError,
		Status: true,
		Data:   err,
	}

	WriteFromJsonBody(writer, webResponse)
}
func WriteFromJsonBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)

	if err != nil {
		panic(err)

	}
}
