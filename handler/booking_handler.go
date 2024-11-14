package handler

import (
	"encoding/json"
	"net/http"

	"github.com/golang-generic/model"
	"github.com/golang-generic/service"
)

type BookingHandler struct {
	service service.BookingService
}

func NewBookingHandler(service service.BookingService) *BookingHandler {
	return &BookingHandler{service}
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking model.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateBooking(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}
