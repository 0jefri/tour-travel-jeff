package dto

import "github.com/golang-generic/model"

type ReviewResponse struct {
	ID          int               `json:"id"`
	Transaction model.Transaction `json:"transaction"`
	Rating      float64
}
