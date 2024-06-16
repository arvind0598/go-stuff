package flights

import "sukasa/bookings/internal/users"

type FlightService interface {
	GetFlightByNumber(flightNumber string) (Flight, bool)
	ReserveSeat(flightNumber string, seatNumber string, passenger Passenger, currentUser users.User) bool
	ResetSeat(flightNumber string, seatNumber string, currentUser users.User) bool
}

type flightService struct {
	repository FlightRepository
}

func (s flightService) GetFlightByNumber(flightNumber string) (Flight, bool) {
	return s.repository.FindFlightByNumber(flightNumber)
}

func (s flightService) ReserveSeat(flightNumber string, seatNumber string, passenger Passenger, currentUser users.User) bool {
	flight, ok := s.repository.FindFlightByNumber(flightNumber)
	if !ok {
		return false
	}

	flight.ReserveSeat(seatNumber, passenger, currentUser.ID)
	return s.repository.UpdateFlightById(flight.ID, flight)
}

func (s flightService) ResetSeat(flightNumber string, seatNumber string, currentUser users.User) bool {
	flight, ok := s.repository.FindFlightByNumber(flightNumber)
	if !ok {
		return false
	}

	flight.ResetSeat(seatNumber, currentUser.ID)
	return s.repository.UpdateFlightById(flight.ID, flight)
}

func GetFlightService() FlightService {
	repository := GetRepository()
	return flightService{repository}
}
