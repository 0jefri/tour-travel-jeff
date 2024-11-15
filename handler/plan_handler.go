package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-generic/service"
)

type PlanHandler struct {
	planService service.PlanService
}

func NewPlanHandler(planService service.PlanService) *PlanHandler {
	return &PlanHandler{planService: planService}
}

func (h *PlanHandler) GetPlansByTourID(w http.ResponseWriter, r *http.Request) {
	tourIDStr := r.URL.Query().Get("tourID")
	if tourIDStr == "" {
		http.Error(w, "TourID is required", http.StatusBadRequest)
		return
	}

	tourID, err := strconv.Atoi(tourIDStr)
	if err != nil {
		http.Error(w, "Invalid TourID", http.StatusBadRequest)
		return
	}

	plans, benefits, err := h.planService.GetPlansByTourID(tourID)
	if err != nil {
		log.Printf("Error retrieving plans: %v", err)
		http.Error(w, "Error retrieving plans", http.StatusInternalServerError)
		return
	}

	response := make([]map[string]interface{}, 0)

	for _, plan := range plans {
		planResponse := map[string]interface{}{
			"id":          plan.ID,
			"tourID":      plan.TourID,
			"title":       plan.Title,
			"description": plan.Description,
			"benefit":     benefits[plan.ID],
		}
		response = append(response, planResponse)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
