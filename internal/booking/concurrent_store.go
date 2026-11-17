// An updated version of MemoryStore that safely handles concurrent read/writes into it by using read/write locks

package booking

import "sync"

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{}
}

func (cs *ConcurrentStore) Book(b Booking) error {
	cs.Lock()
	defer cs.Unlock()

	if _, exists := cs.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}

	cs.bookings[b.SeatID] = b

	return nil
}

func (cs *ConcurrentStore) ListBookings(movieID string) []Booking {
	cs.RLock()
	defer cs.RUnlock()

	var bookings []Booking

	for _, b := range cs.bookings {
		if b.MovieID == movieID {
			bookings = append(bookings, b)
		}
	}

	return bookings
}
