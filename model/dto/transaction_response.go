package dto

import "github.com/golang-generic/model"

type TransactionResponse struct {
	ID     int        `json:"id"`
	Tour   model.Tour `json:"tour"`
	Status string     `json:"status"`
}
