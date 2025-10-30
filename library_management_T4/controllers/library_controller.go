package controllers

import (
	"fmt"
	"library_management_T4/services"
	"time"
)

type LibraryController struct {
	service services.LibraryManager
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{service: service}
}

func (c *LibraryController) ReserveBook(bookID, memberID int) {
	fmt.Printf("[REQUEST] Member %d trying to reserve Book %d...\n", memberID, bookID)
	err := c.service.ReserveBook(bookID, memberID)
	if err != nil {
		fmt.Printf("[FAILED] Member %d: %v\n", memberID, err)
	} else {
		fmt.Printf("[SUCCESS] Member %d reserved Book %d\n", memberID, bookID)
	}
}

func (c *LibraryController) ListBooks() {
	books := c.service.GetBooks()
	fmt.Println("\n--- Current Books ---")
	for _, b := range books {
		status := "Available"
		if !b.Available {
			status = fmt.Sprintf("Reserved by Member %d", b.ReservedBy)
		}
		fmt.Printf("ID: %d | %s | %s\n", b.ID, b.Title, status)
	}
	fmt.Println()
}

func (c *LibraryController) SimulateConcurrentReservations() {
	fmt.Println("Simulating concurrent reservation attempts...\n")
	time.Sleep(100 * time.Millisecond)

	// Simulate 10 members trying to reserve the same book
	for i := 0; i < 10; i++ {
		memberID := 101 + (i % 3) // cycle through Alice, Bob, Charlie
		go c.ReserveBook(1, memberID)
	}

	// Wait a bit
	time.Sleep(2 * time.Second)
	c.ListBooks()

	// Wait for timeout
	fmt.Println("Waiting for auto-cancellation (5 seconds)...")
	time.Sleep(6 * time.Second)
	c.ListBooks()
}
