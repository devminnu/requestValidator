package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Booking(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"hello"}`))
}
