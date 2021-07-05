package main

import (
	"log"
	"net/http"
	"time"

	"requestvalidator/internal/app/handler"
	"requestvalidator/middleware"
	"requestvalidator/rules"

	"github.com/gorilla/mux"
)

func main() {
	rules.BookingRules["checkIn"] = []string{"required", "date"}
	delete(rules.BookingRules, "bookingId")

	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.Use(middleware.RequestValidation)
	r.HandleFunc("/booking", handler.Booking)

	log.Fatal(srv.ListenAndServe())
}
