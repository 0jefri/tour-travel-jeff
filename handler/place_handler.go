package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-generic/service"
)

type PlaceHandler struct {
	placeService service.PlaceService
}

func NewPlaceHandler(placeService service.PlaceService) *PlaceHandler {
	return &PlaceHandler{placeService: placeService}
}

func (h *PlaceHandler) GetAllPlaces(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	pageParam := r.URL.Query().Get("page")
	sortParam := r.URL.Query().Get("sort")
	filterParam := r.URL.Query().Get("filter")
	dateParam := r.URL.Query().Get("date")

	limit := 10
	page := 1

	if limitParam != "" {
		parsedLimit, err := strconv.Atoi(limitParam)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if pageParam != "" {
		parsedPage, err := strconv.Atoi(pageParam)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if sortParam == "" {
		sortParam = "low-to-high"
	}

	if filterParam == "" {
		filterParam = "all"
	}

	if dateParam != "" {
		_, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}
	}

	places, err := h.placeService.GetAllPlaces(limit, page, sortParam, filterParam, dateParam)
	if err != nil {
		http.Error(w, "Error fetching places", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(places); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
