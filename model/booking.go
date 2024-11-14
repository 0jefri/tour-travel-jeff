package model

import "time"

type Booking struct {
	ID             int
	Name           string
	Email          string
	ConfirmEmail   string
	Phone          string
	Date           time.Time
	NumberOfTicket string
	Message        string
	Tour           Tour
}
