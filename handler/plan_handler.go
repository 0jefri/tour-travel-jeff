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

type ResponseWrapper struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (h *PlanHandler) GetPlansByTourID(w http.ResponseWriter, r *http.Request) {
	tourIDStr := r.URL.Query().Get("tourID")
	if tourIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseWrapper{
			Status:  http.StatusBadRequest,
			Message: "TourID is required",
			Data:    nil,
		})
		return
	}

	tourID, err := strconv.Atoi(tourIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseWrapper{
			Status:  http.StatusBadRequest,
			Message: "Invalid TourID",
			Data:    nil,
		})
		return
	}

	plans, benefits, err := h.planService.GetPlansByTourID(tourID)
	if err != nil {
		log.Printf("Error retrieving plans: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseWrapper{
			Status:  http.StatusInternalServerError,
			Message: "Error retrieving plans",
			Data:    nil,
		})
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
	json.NewEncoder(w).Encode(ResponseWrapper{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	})
}
