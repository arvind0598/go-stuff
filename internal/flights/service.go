package flights

import (
	"errors"
	"sukasa/bookings/internal/users"
	"time"
)

type FlightService interface {
	GetFlightByNumber(flightNumber string) (Flight, bool)
	ReserveSeat(flightNumber string, seatNumber string, passenger Passenger, currentUser users.User) error
	ResetSeat(flightNumber string, seatNumber string, currentUser users.User) error
}

type flightService struct {
	repository FlightRepository
}

func (s flightService) GetFlightByNumber(flightNumber string) (Flight, bool) {
	return s.repository.FindFlightByNumber(flightNumber)
}

func (s flightService) findSeat(flightNumber string, seatNumber string) (*Flight, int, error) {
	flight, ok := s.repository.FindFlightByNumber(flightNumber)
	if !ok {
		return nil, -1, errors.New("flight not found")
	}

	for i, seat := range flight.Seats {
		if seat.SeatNumber == seatNumber {
			return &flight, i, nil
		}
	}

	return nil, -1, errors.New("seat not found")
}

func (s flightService) ReserveSeat(flightNumber string, seatNumber string, passenger Passenger, currentUser users.User) error {
	flight, seatIndex, err := s.findSeat(flightNumber, seatNumber)
	if err != nil {
		return err
	}

	if flight.Seats[seatIndex].Passenger.Phone != "" {
		return errors.New("seat already reserved")
	}

	changeLog := BookingChangeLog{currentUser.ID, "", passenger.Phone, time.Now().String()}
	flight.Seats[seatIndex].ChangeLogs = append(flight.Seats[seatIndex].ChangeLogs, changeLog)
	flight.Seats[seatIndex].Passenger = passenger
	if !s.repository.UpdateFlightById(flight.ID, *flight) {
		return errors.New("failed to update flight")
	}

	return nil
}

func (s flightService) ResetSeat(flightNumber string, seatNumber string, currentUser users.User) error {
	flight, seatIndex, err := s.findSeat(flightNumber, seatNumber)
	if err != nil {
		return err
	}

	if flight.Seats[seatIndex].Passenger.Phone == "" {
		return errors.New("seat not reserved")
	}

	changeLog := BookingChangeLog{currentUser.ID, flight.Seats[seatIndex].Passenger.Phone, "", time.Now().String()}
	flight.Seats[seatIndex].ChangeLogs = append(flight.Seats[seatIndex].ChangeLogs, changeLog)
	flight.Seats[seatIndex].Passenger = Passenger{}
	if !s.repository.UpdateFlightById(flight.ID, *flight) {
		return errors.New("failed to update flight")
	}

	return nil
}

func GetFlightService() FlightService {
	repository := GetRepository()
	return flightService{repository}
}
