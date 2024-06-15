package main

import "fmt"

type Passenger struct {
	name    string
	contact string
	age     string
}

type Seat struct {
	seatNumber string
	passenger  *Passenger
}

func (s Seat) IsReserved() bool {
	return s.passenger != nil
}

func (s *Seat) Reserve(p *Passenger) {
	s.passenger = p
}

func (s *Seat) Reset() {
	s.passenger = nil
}

var row_seats = [...]string{"A", "B", "C", "D", "E", "F"}

const NUM_ROWS = 10

var seats = map[string]*Seat{}

func initSeats() {
	for i := 1; i <= NUM_ROWS; i++ {
		for _, seat := range row_seats {
			key := fmt.Sprintf("%d%s", i, seat)
			seats[key] = &Seat{seatNumber: key}
		}
	}
}
