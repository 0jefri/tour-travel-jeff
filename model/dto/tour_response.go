package dto

import (
	"time"

	"github.com/golang-generic/model"
)

type TourResponse struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Place model.Place `json:"place"`
	Date  time.Time   `json:"date"`
}
