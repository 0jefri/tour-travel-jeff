package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-generic/database"
	"github.com/golang-generic/handler"
	"github.com/golang-generic/repository"
	"github.com/golang-generic/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Println("Successfully connected to the database!")
	}
	defer db.Close()

	placeRepository := repository.NewPlaceRepository(db)
	placeService := service.NewPlaceService(placeRepository)
	placeHandler := handler.NewPlaceHandler(placeService)

	bookingRepository := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepository)
	bookingHandler := handler.NewBookingHandler(bookingService)

	plansRepository := repository.NewPlanRepository(db)
	plansService := service.NewPlanService(plansRepository)
	plansHandler := handler.NewPlanHandler(plansService)

	locationRepository := repository.NewLocationRepository(db)
	locationService := service.NewLocationService(locationRepository)
	locationHandler := handler.NewLocationHandler(locationService)

	r := mux.NewRouter()
	r.HandleFunc("/places", placeHandler.GetAllPlaces).Methods("GET")
	r.HandleFunc("/places-detail", placeHandler.GetPlaceDetail).Methods("GET")
	r.HandleFunc("/booking", bookingHandler.CreateBooking).Methods("POST")
	r.HandleFunc("/plans", plansHandler.GetPlansByTourID).Methods("GET")
	r.HandleFunc("/location", locationHandler.GetLocationsByTourID).Methods("GET")
	// r.HandleFunc("/location/{id}", locationHandler.GetLocationByID).Methods("GET")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
