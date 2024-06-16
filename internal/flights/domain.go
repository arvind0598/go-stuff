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
	SeatNumber       string             `bson:"seatNumber"`
	Passenger        Passenger          `bson:"passenger"`
	ReservedByUserId primitive.ObjectID `bson:"reservedByUserId"`
	ResetLogs        []BookingResetLog  `bson:"resetLogs"`
}

type BookingResetLog struct {
	ResetByUser            primitive.ObjectID `bson:"resetByUser"`
	PreviousPassengerPhone string             `bson:"previousPassengerPhone"`
	ResetAt                string             `bson:"resetAt"`
}

type Passenger struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
	Age   int    `bson:"age"`
}
