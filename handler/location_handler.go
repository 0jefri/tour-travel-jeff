package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-generic/service"
)

type LocationHandler struct {
	locationService service.LocationService
}

func NewLocationHandler(locationService service.LocationService) *LocationHandler {
	return &LocationHandler{locationService: locationService}
}

func (h *LocationHandler) GetLocationsByTourID(w http.ResponseWriter, r *http.Request) {
	tourIDStr := r.URL.Query().Get("tourID")
	if tourIDStr == "" {
		http.Error(w, "tour_id is required", http.StatusBadRequest)
		return
	}

	tourID, err := strconv.Atoi(tourIDStr)
	if err != nil {
		http.Error(w, "Invalid tour_id", http.StatusBadRequest)
		return
	}

	locations, err := h.locationService.GetLocationsByTourID(tourID)
	if err != nil {
		log.Println("error:", err)
		log.Printf("Error retrieving locations: %v", err)
		http.Error(w, "Error retrieving locations", http.StatusInternalServerError)
		return
	}

	if len(locations) == 0 {
		http.Error(w, "No locations found for this tour", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"success":     true,
		"status_code": http.StatusOK,
		"message":     "Location Found",
		"data":        locations,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
