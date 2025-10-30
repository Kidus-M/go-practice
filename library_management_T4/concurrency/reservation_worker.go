package concurrency

import (
	"library_management_T4/models"
	"log"
	"time"
)

type ReservationRequest struct {
	BookID   int
	MemberID int
	Member   *models.Member
	Book     *models.Book
	Response chan error
}

// Interface expected by the worker pool
type ReservationHandler interface {
	TryReserve(bookID, memberID int) error
	CancelReservation(bookID int)
}

type ReservationWorkerPool struct {
	Requests chan *ReservationRequest
	handler  ReservationHandler
	quit     chan bool
}

// NewReservationWorkerPool creates a worker pool with a handler
func NewReservationWorkerPool(handler ReservationHandler) *ReservationWorkerPool {
	return &ReservationWorkerPool{
		Requests: make(chan *ReservationRequest, 100),
		handler:  handler,
		quit:     make(chan bool),
	}
}

// Start launches worker goroutines
func (p *ReservationWorkerPool) Start() {
	for i := 0; i < 5; i++ {
		go p.worker()
	}
}

// worker processes incoming requests
func (p *ReservationWorkerPool) worker() {
	for {
		select {
		case req := <-p.Requests:
			p.process(req)
		case <-p.quit:
			return
		}
	}
}

// process handles one reservation
func (p *ReservationWorkerPool) process(req *ReservationRequest) {
	err := p.handler.TryReserve(req.BookID, req.MemberID)
	if err != nil {
		req.Response <- err
		return
	}

	log.Printf("Book %d reserved by Member %d (%s)\n", req.BookID, req.MemberID, req.Member.Name)

	// Auto-cancel after 5 seconds
	timer := time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C
		log.Printf("Reservation timeout – releasing Book %d\n", req.BookID)
		p.handler.CancelReservation(req.BookID)
	}()

	req.Response <- nil
}

// Stop shuts down the pool
func (p *ReservationWorkerPool) Stop() {
	close(p.quit)
}
