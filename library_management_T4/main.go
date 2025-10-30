package main

import (
	"library_management_T4/controllers"
	"library_management_T4/services"
)

func main() {
	service := services.NewLibraryService()
	controller := controllers.NewLibraryController(service)

	// Initial state
	controller.ListBooks()

	// Simulate concurrent reservations
	controller.SimulateConcurrentReservations()

	// Keep main alive
	select {}
}
