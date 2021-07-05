package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"requestvalidator/rules"

	"github.com/thedevsaddam/govalidator"
)

func RequestValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		data := make(map[string]interface{}, 0)

		opts := govalidator.Options{
			Request: r,
			Data:    &data,
		}
		switch r.RequestURI {
		case "/booking":
			opts.Rules = rules.BookingRules
		}

		vd := govalidator.New(opts)
		e := vd.ValidateJSON()
		w.Header().Set("Content-type", "application/json")
		if len(e) != 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
			return
		}
		next.ServeHTTP(w, r)
	})

}
