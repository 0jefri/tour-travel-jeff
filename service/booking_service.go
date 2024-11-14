package service

import (
	"errors"
	"net/mail"

	"github.com/golang-generic/model"
	"github.com/golang-generic/repository"
)

type BookingService interface {
	CreateBooking(booking *model.Booking) error
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{repo}
}

func (s *bookingService) CreateBooking(booking *model.Booking) error {
	if _, err := mail.ParseAddress(booking.Email); err != nil {
		return errors.New("invalid email format")
	}

	if booking.Email != booking.ConfirmEmail {
		return errors.New("email confirmation failed")
	}

	// bookingResponse := model.Booking{
	// 	ID: booking.ID,
	// 	Name: booking.Name,
	// 	Email: booking.Email,
	// 	ConfirmEmail: booking.ConfirmEmail,
	// 	Phone: booking.Phone,
	// 	Date: time.Date(),
	// 	Message: booking.Message,
	// 	Tour: model.Tour{
	// 		ID: booking.Tour.ID,
	// 		Name: booking.Tour.Name,
	// 		Date: time.Date(),
	// 	},
	// }

	return s.repo.CreateBooking(booking)
}
