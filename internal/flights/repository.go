package flights

import (
	"context"
	"sukasa/bookings/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepository interface {
	FindFlightByNumber(flightNumber string) (Flight, bool)
	FindFlightById(id primitive.ObjectID) Flight
	UpdateFlightById(id primitive.ObjectID, flight Flight) bool
}

type flightRepository struct {
	client *mongo.Client
}

func (r flightRepository) FindFlightByNumber(flightNumber string) (Flight, bool) {
	collection := r.client.Database("sukasa").Collection("flights")
	filter := bson.M{"flightNumber": flightNumber}

	var result Flight
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return Flight{}, false
	}

	return result, true
}

func (r flightRepository) FindFlightById(id primitive.ObjectID) Flight {
	collection := r.client.Database("sukasa").Collection("flights")
	filter := bson.M{"_id": id}

	var result Flight
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return Flight{}
	}

	return result
}

func (r flightRepository) UpdateFlightById(id primitive.ObjectID, flight Flight) bool {
	collection := r.client.Database("sukasa").Collection("flights")
	filter := bson.M{"_id": id}

	_, err := collection.ReplaceOne(context.TODO(), filter, flight)
	return err == nil
}

func GetRepository() FlightRepository {
	client := db.GetClient()
	return flightRepository{client}
}
