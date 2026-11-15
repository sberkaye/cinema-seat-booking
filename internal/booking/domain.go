package booking

import "errors"

// Common error types
var (
	ErrSeatAlreadyBooked = errors.New("This seat is already taken")
)

// Booking struct represents a confirmed seat booking.
type Booking struct {
	ID      string
	MovieID string
	SeatID  string
	UserID  string
	Status  string
}

// API expected to be implemented by the booking stores
type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}
