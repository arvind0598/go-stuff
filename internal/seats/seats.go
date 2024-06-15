package seats

import (
	"fmt"
	"sukasa/bookings/internal/passengers"
)

type Seat struct {
	seatNumber string
	passenger  *passengers.Passenger
}

func (s Seat) IsReserved() bool {
	return s.passenger != nil
}

func (s *Seat) Reserve(p *passengers.Passenger) {
	s.passenger = p
}

func (s *Seat) Reset() {
	s.passenger = nil
}

var seats = map[string]*Seat{}

func CreateSeats() {
	var rowSeats = [...]string{"A", "B", "C", "D", "E", "F"}
	var numberOfRows = 10

	for i := 1; i <= numberOfRows; i++ {
		for _, seat := range rowSeats {
			key := fmt.Sprintf("%d%s", i, seat)
			seats[key] = &Seat{seatNumber: key}
		}
	}
}
