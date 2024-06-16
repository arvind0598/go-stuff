package flights

import "go.mongodb.org/mongo-driver/bson/primitive"

type Flight struct {
	ID               primitive.ObjectID `bson:"_id"`
	FlightNumber     string             `bson:"flightNumber"`
	DepartureAirport string             `bson:"departureAirport"`
	ArrivalAirport   string             `bson:"arrivalAirport"`
	Seats            []Seat             `bson:"seats"`
}

type Seat struct {
	SeatNumber string             `bson:"seatNumber"`
	Passenger  Passenger          `bson:"passenger"`
	ChangeLogs []BookingChangeLog `bson:"changeLogs"`
}

type BookingChangeLog struct {
	ChangedByUserId       primitive.ObjectID `bson:"changedByUserId"`
	InitialPassengerPhone string             `bson:"initialPassengerPhone"`
	NewPassengerPhone     string             `bson:"newPassengerPhone"`
	ChangeAt              string             `bson:"changeAt"`
}

type Passenger struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
	Age   int    `bson:"age"`
}
