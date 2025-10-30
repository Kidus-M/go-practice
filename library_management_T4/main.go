package main

import (
	"library_management_T4/controllers"
	"library_management_T4/services"
	"log"
	"time"
)

func main() {
	service := services.NewLibraryService()
	controller := controllers.NewLibraryController(service)

	// Initial state
	controller.ListBooks()

	// Simulate concurrent reservations
	controller.SimulateConcurrentReservations()

	// Let final logs print
	time.Sleep(1 * time.Second)

	// Gracefully stop worker pool
	if service.WorkerPool != nil {
		service.WorkerPool.Stop()
	}

	log.Println("Simulation complete. Shutting down.")
}
