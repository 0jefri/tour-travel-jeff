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
