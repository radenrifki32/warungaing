package helper

import (
	"encoding/json"
	"net/http"

	"github.com/rifki321/warungku/user/web"
)

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	var response interface{}
	status := true
	if code >= 400 {
		status = false
	}
	if payload != nil {
		response = &web.ResponseWeb{
			Status: status,
			Code:   code,
			Data:   payload,
		}
	} else {
		response = &web.ResponseWebWithMessage{
			Message: message,
			Status:  status,
			Code:    code,
		}
	}

	res, _ := json.Marshal(&response)
	w.Write(res)
}
