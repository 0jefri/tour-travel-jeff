package service

import (
	"github.com/golang-generic/model"
	"github.com/golang-generic/repository"
)

type PlaceService interface {
	GetAllPlaces(limit, page int, sort, filter, date string) ([]repository.PlaceWithDate, error)
	GetPlaceDetail(id int) (*model.PlaceDetail, error)
}

type placeService struct {
	placeRepo repository.PlaceRepository
}

func NewPlaceService(placeRepo repository.PlaceRepository) PlaceService {
	return &placeService{placeRepo}
}

func (s *placeService) GetAllPlaces(limit, page int, sort, filter, date string) ([]repository.PlaceWithDate, error) {
	return s.placeRepo.GetAllPlaces(limit, page, sort, filter, date)
}

func (s *placeService) GetPlaceDetail(placeDetailID int) (*model.PlaceDetail, error) {
	return s.placeRepo.GetPlaceDetail(placeDetailID)
}
