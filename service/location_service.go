package service

import (
	"github.com/golang-generic/model/dto"
	"github.com/golang-generic/repository"
)

type LocationService interface {
	GetLocationsByTourID(tourID int) ([]dto.LocationResponse, error)
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) GetLocationsByTourID(tourID int) ([]dto.LocationResponse, error) {
	locations, err := s.repo.GetLocationsByTourID(tourID)
	if err != nil {
		return nil, err
	}

	var responses []dto.LocationResponse
	for _, loc := range locations {
		responses = append(responses, dto.LocationResponse{
			ID:          loc.ID,
			Title:       loc.Title,
			MapUrl:      loc.MapUrl,
			Description: loc.Description,
			TourID:      loc.TourID,
		})
	}

	return responses, nil
}
