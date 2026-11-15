// Naive in-memory BookingStore implementation using the map data structure

package booking

type MemoryStore struct {
	bookings map[string]Booking // store bookings mapped to the seat IDs they are booked for
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Book(b Booking) error {
	// check if the booking already exists, return with an error if so
	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b

	return nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	bookings := make([]Booking, len(s.bookings))

	for _, b := range s.bookings {
		if b.MovieID == movieID {
			bookings = append(bookings, b)
		}
	}

	return bookings
}
