package dto

type LocationResponse struct {
	ID          int    `json:"id"`
	TourID      int    `json:"tourID"`
	Title       string `json:"title"`
	MapUrl      string `json:"mapUrl"`
	Description string `json:"description"`
}
