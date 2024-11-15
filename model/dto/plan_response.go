package dto

import "github.com/golang-generic/model"

type PlanResponse struct {
	ID          int             `json:"id"`
	TourID      int             `json:"tourID"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Benefit     []model.Benefit `json:"benefit"`
}
