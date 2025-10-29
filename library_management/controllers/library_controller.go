package services

import (
	"errors"
	"sync"

	"library_management/models"
)

type Library struct {
	books   map[int]models.Book
	members map[int]*models.Member
	mu      sync.Mutex
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]*models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if book.Status == "" {
		book.Status = "Available"
	}
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.books, bookID)
}

func (l *Library) AddMember(id int, name string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, exists := l.members[id]; !exists {
		l.members[id] = &models.Member{ID: id, Name: name}
	}
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}

	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found")
	}

	found := false
	idx := -1
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			found = true
			idx = i
			break
		}
	}
	if !found {
		return errors.New("member did not borrow this book")
	}

	member.BorrowedBooks = append(member.BorrowedBooks[:idx], member.BorrowedBooks[idx+1:]...)
	book.Status = "Available"
	l.books[bookID] = book
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	var available []models.Book
	for _, b := range l.books {
		if b.Status == "Available" {
			available = append(available, b)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	member, ok := l.members[memberID]
	if !ok {
		return []models.Book{}
	}
	return member.BorrowedBooks
}
