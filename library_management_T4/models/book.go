package models

type Book struct {
	ID         int
	Title      string
	Available  bool
	ReservedBy int
}
