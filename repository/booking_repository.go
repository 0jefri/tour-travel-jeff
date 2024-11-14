package repository

import (
	"database/sql"

	"github.com/golang-generic/model"
)

type BookingRepository interface {
	CreateBooking(booking *model.Booking) error
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking *model.Booking) error {
	query := `INSERT INTO bookings (name, email, confirm_email, phone, date, number_of_ticket, message, tour_id) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := r.db.QueryRow(query, booking.Name, booking.Email, booking.ConfirmEmail, booking.Phone, booking.Date, booking.NumberOfTicket, booking.Message, booking.Tour.ID).Scan(&booking.ID)
	return err
}
