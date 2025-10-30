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

	controller.ListBooks()

	controller.SimulateConcurrentReservations()

	time.Sleep(1 * time.Second)

	if service.WorkerPool != nil {
		service.WorkerPool.Stop()
	}

	log.Println("Simulation complete. Shutting down.")
}
