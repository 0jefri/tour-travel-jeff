package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-generic/model"
)

type PlaceWithDate struct {
	Place model.Place
	Date  time.Time
}

type PlaceRepository interface {
	GetAllPlaces(limit, page int, sort, filter, date string) ([]PlaceWithDate, error)
	GetPlaceDetail(placeDetailID int) (*model.PlaceDetail, error)
}

type placeRepository struct {
	db *sql.DB
}

func NewPlaceRepository(db *sql.DB) PlaceRepository {
	return &placeRepository{db}
}

func (r *placeRepository) GetAllPlaces(limit, page int, sort, filter, date string) ([]PlaceWithDate, error) {
	offset := (page - 1) * limit

	orderBy := ""
	if sort == "low-to-high" {
		orderBy = "ORDER BY p.price ASC"
	} else if sort == "high-to-low" {
		orderBy = "ORDER BY p.price DESC"
	}

	dateFilter := ""
	if date != "" {
		dateFilter = fmt.Sprintf("AND t.date::date = '%s'", date)
	}

	filterQuery := ""
	if filter == "all" {
		filterQuery = ""
	} else {
		filterQuery = "WHERE p.name IS NOT NULL"
	}

	query := fmt.Sprintf(`SELECT p.id, p.name, p.description, p.photo, p.price, t.date 
                          FROM Place p 
                          LEFT JOIN Tour t ON p.id = t.place_id
                          %s
                          %s
                          %s
                          LIMIT %d OFFSET %d`, filterQuery, dateFilter, orderBy, limit, offset)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []PlaceWithDate
	for rows.Next() {
		var place model.Place
		var date *time.Time
		if err := rows.Scan(&place.ID, &place.Name, &place.Description, &place.Photo, &place.Price, &date); err != nil {

			fmt.Println("Scanned date:", date)
			return nil, err
		}
		var placeWithDate PlaceWithDate
		placeWithDate.Place = place
		if date != nil {
			placeWithDate.Date = *date
		}
		places = append(places, placeWithDate)
	}

	return places, nil
}

func (r *placeRepository) GetPlaceDetail(placeDetailID int) (*model.PlaceDetail, error) {
	// Query untuk mengambil place detail
	query := `SELECT p.id, p.name, r.id, r.rating
						FROM place_details pd
						JOIN place p ON pd.place_id = p.id
						JOIN review r ON pd.review_id = r.id
						WHERE pd.id = $1`

	placeDetail := &model.PlaceDetail{}
	err := r.db.QueryRow(query, placeDetailID).Scan(&placeDetail.Place.ID, &placeDetail.Place.Name, &placeDetail.Review.ID, &placeDetail.Review.Rating)
	if err != nil {
		return nil, fmt.Errorf("error fetching place detail: %v", err)
	}

	photoQuery := `SELECT id, url, caption
								 FROM photos
								 WHERE photo_group_id = (SELECT id FROM photo_groups WHERE place_detail_id = $1) LIMIT 4`

	rows, err := r.db.Query(photoQuery, placeDetailID)
	if err != nil {
		return nil, fmt.Errorf("error fetching photos: %v", err)
	}
	defer rows.Close()

	var photos []model.Photo
	for rows.Next() {
		var photo model.Photo
		err := rows.Scan(&photo.ID, &photo.URL, &photo.Caption)
		if err != nil {
			return nil, fmt.Errorf("error scanning photo: %v", err)
		}
		photos = append(photos, photo)
	}
	// placeDetail.Galery = photos
	placeDetail.Galery = append(placeDetail.Galery, model.PhotoGroup{Photos: photos})

	return placeDetail, nil
}
