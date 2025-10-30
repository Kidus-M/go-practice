package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

type Controller struct {
	lib     services.LibraryManager
	scanner *bufio.Scanner
}

func NewController(lib services.LibraryManager) *Controller {
	return &Controller{
		lib:     lib,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *Controller) Run() {
	c.seedData()

	for {
		c.printMenu()
		choice := c.readInt("Enter choice: ")
		switch choice {
		case 1:
			c.addBook()
		case 2:
			c.removeBook()
		case 3:
			c.borrowBook()
		case 4:
			c.returnBook()
		case 5:
			c.listAvailable()
		case 6:
			c.listBorrowed()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (c *Controller) seedData() {
	c.lib.AddBook(models.Book{ID: 1, Title: "Go Lang", Author: "Alan"})
	c.lib.AddBook(models.Book{ID: 2, Title: "Clean Code", Author: "Uncle Bob"})
	c.lib.(*services.Library).AddMember(models.Member{ID: 101, Name: "Alice"})
	c.lib.(*services.Library).AddMember(models.Member{ID: 102, Name: "Bob"})
}

func (c *Controller) printMenu() {
	fmt.Println("\n=== Library System ===")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available")
	fmt.Println("6. List Borrowed by Member")
	fmt.Println("7. Exit")
}

func (c *Controller) addBook() {
	id := c.readInt("Book ID: ")
	title := c.readString("Title: ")
	author := c.readString("Author: ")
	c.lib.AddBook(models.Book{ID: id, Title: title, Author: author})
	fmt.Println("Book added.")
}

func (c *Controller) removeBook() {
	id := c.readInt("Book ID to remove: ")
	c.lib.RemoveBook(id)
	fmt.Println("Book removed if existed.")
}

func (c *Controller) borrowBook() {
	bookID := c.readInt("Book ID: ")
	memberID := c.readInt("Member ID: ")
	if err := c.lib.BorrowBook(bookID, memberID); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Borrowed!")
	}
}

func (c *Controller) returnBook() {
	bookID := c.readInt("Book ID: ")
	memberID := c.readInt("Member ID: ")
	if err := c.lib.ReturnBook(bookID, memberID); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Returned!")
	}
}

func (c *Controller) listAvailable() {
	books := c.lib.ListAvailableBooks()
	fmt.Println("\n--- Available ---")
	for _, b := range books {
		fmt.Printf("%d | %s by %s\n", b.ID, b.Title, b.Author)
	}
	if len(books) == 0 {
		fmt.Println("None")
	}
}

func (c *Controller) listBorrowed() {
	id := c.readInt("Member ID: ")
	books := c.lib.ListBorrowedBooks(id)
	m, _ := c.lib.(*services.Library).GetMember(id)
	name := "Unknown"
	if m.ID != 0 {
		name = m.Name
	}
	fmt.Printf("\n--- Borrowed by %s (%d) ---\n", name, id)
	for _, b := range books {
		fmt.Printf("%d | %s by %s\n", b.ID, b.Title, b.Author)
	}
	if len(books) == 0 {
		fmt.Println("None")
	}
}

// Input helpers
func (c *Controller) readString(prompt string) string {
	fmt.Print(prompt)
	c.scanner.Scan()
	return strings.TrimSpace(c.scanner.Text())
}

func (c *Controller) readInt(prompt string) int {
	for {
		s := c.readString(prompt)
		if n, err := strconv.Atoi(s); err == nil {
			return n
		}
		fmt.Println("Invalid number, try again.")
	}
}
