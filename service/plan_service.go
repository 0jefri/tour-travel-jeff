package service

import (
	"log"

	"github.com/golang-generic/model"
	"github.com/golang-generic/repository"
)

type PlanService interface {
	GetPlansByTourID(tourID int) ([]model.Plan, map[int][]model.Benefit, error)
}

type planService struct {
	planRepository repository.PlanRepository
}

func NewPlanService(planRepository repository.PlanRepository) PlanService {
	return &planService{planRepository: planRepository}
}

func (s *planService) GetPlansByTourID(tourID int) ([]model.Plan, map[int][]model.Benefit, error) {
	plans, benefits, err := s.planRepository.FindPlansByTourID(tourID)
	if err != nil {
		log.Printf("Error retrieving plans: %v", err)
		return nil, nil, err
	}

	return plans, benefits, nil
}
