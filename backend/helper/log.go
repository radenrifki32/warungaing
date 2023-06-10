package helper

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func LogBody(w http.ResponseWriter, r *http.Request) {
	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
		return
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	log.Printf("BODY: %q", rdr1)
	r.Body = rdr2
}
