package flights

import "go.mongodb.org/mongo-driver/bson/primitive"

type Flight struct {
	ID               primitive.ObjectID `bson:"_id"`
	FlightNumber     string
	DepartureAirport string
	ArrivalAirport   string
	Seats            []Seat
}

func (f *Flight) ReserveSeat(seatNumber string, passenger Passenger, userId primitive.ObjectID) {
	for i, seat := range f.Seats {
		if seat.SeatNumber == seatNumber {
			f.Seats[i].Passenger = passenger
			f.Seats[i].ReservedByUserId = userId
		}
	}
}

func (f *Flight) ResetSeat(seatNumber string, userId primitive.ObjectID) {
	for i, seat := range f.Seats {
		if seat.SeatNumber == seatNumber {
			previousPassengerPhone := f.Seats[i].Passenger.Phone
			f.Seats[i].Passenger = Passenger{}
			f.Seats[i].ReservedByUserId = primitive.NilObjectID
			f.Seats[i].ResetLogs = append(f.Seats[i].ResetLogs, BookingResetLog{
				ResetByUser:            userId,
				PreviousPassengerPhone: previousPassengerPhone,
			})
		}
	}
}

type Seat struct {
	SeatNumber       string
	Passenger        Passenger
	ReservedByUserId primitive.ObjectID `bson:"reservedByUser"`
	ResetLogs        []BookingResetLog
}

type BookingResetLog struct {
	ResetByUser            primitive.ObjectID `bson:"resetByUser"`
	PreviousPassengerPhone string
	ResetAt                string
}

type Passenger struct {
	Name  string
	Phone string
	Age   int
}
