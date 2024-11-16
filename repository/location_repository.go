package repository

import (
	"database/sql"
	"log"

	"github.com/golang-generic/model"
)

type LocationRepository interface {
	GetLocationsByTourID(tourID int) ([]model.Location, error)
}

type locationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) GetLocationsByTourID(tourID int) ([]model.Location, error) {
	var locations []model.Location
	query := `SELECT id, title, map_url, description, tour_id FROM locations WHERE tour_id = $1`

	rows, err := r.db.Query(query, tourID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var loc model.Location
		err := rows.Scan(&loc.ID, &loc.Title, &loc.MapUrl, &loc.Description, &loc.TourID)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		locations = append(locations, loc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}
