package rules

import "github.com/thedevsaddam/govalidator"

var (
	BookingRules = govalidator.MapData{
		"bookingId": []string{"required"},
	}
)
