package services

import (
	"errors"
	"library_management_T4/concurrency"
	"library_management_T4/models"
	"sync"
)

var (
	ErrBookNotAvailable = errors.New("book is already reserved")
	ErrBookNotFound     = errors.New("book not found")
	ErrMemberNotFound   = errors.New("member not found")
)

type LibraryManager interface {
	ReserveBook(bookID, memberID int) error
	GetBooks() []models.Book
	GetMembers() []models.Member
}

type ReservationHandler interface {
	TryReserve(bookID, memberID int) error
	CancelReservation(bookID int)
}

type LibraryService struct {
	books      map[int]*models.Book
	members    map[int]*models.Member
	mu         sync.RWMutex
	WorkerPool *concurrency.ReservationWorkerPool
}

func NewLibraryService() *LibraryService {
	books := map[int]*models.Book{
		1: {ID: 1, Title: "The Go Programming Language", Available: true},
		2: {ID: 2, Title: "Concurrency in Go", Available: true},
		3: {ID: 3, Title: "Clean Code", Available: true},
	}

	members := map[int]*models.Member{
		101: {ID: 101, Name: "Alice"},
		102: {ID: 102, Name: "Bob"},
		103: {ID: 103, Name: "Charlie"},
	}

	ls := &LibraryService{
		books:   books,
		members: members,
	}

	ls.WorkerPool = concurrency.NewReservationWorkerPool(ls)
	ls.WorkerPool.Start()

	return ls
}

func (ls *LibraryService) ReserveBook(bookID, memberID int) error {
	ls.mu.RLock()
	book, bookExists := ls.books[bookID]
	member, memberExists := ls.members[memberID]
	ls.mu.RUnlock()

	if !bookExists {
		return ErrBookNotFound
	}
	if !memberExists {
		return ErrMemberNotFound
	}

	req := &concurrency.ReservationRequest{
		BookID:   bookID,
		MemberID: memberID,
		Member:   member,
		Book:     book,
		Response: make(chan error, 1),
	}

	ls.WorkerPool.Requests <- req
	return <-req.Response
}

func (ls *LibraryService) GetBooks() []models.Book {
	ls.mu.RLock()
	defer ls.mu.RUnlock()
	var list []models.Book
	for _, b := range ls.books {
		list = append(list, *b)
	}
	return list
}

func (ls *LibraryService) GetMembers() []models.Member {
	ls.mu.RLock()
	defer ls.mu.RUnlock()
	var list []models.Member
	for _, m := range ls.members {
		list = append(list, *m)
	}
	return list
}

func (ls *LibraryService) TryReserve(bookID, memberID int) error {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	book := ls.books[bookID]
	if book == nil {
		return ErrBookNotFound
	}
	if !book.Available || book.ReservedBy != 0 {
		return ErrBookNotAvailable
	}

	book.Available = false
	book.ReservedBy = memberID
	return nil
}

func (ls *LibraryService) CancelReservation(bookID int) {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	book := ls.books[bookID]
	if book != nil && !book.Available && book.ReservedBy != 0 {
		book.Available = true
		book.ReservedBy = 0
	}
}
