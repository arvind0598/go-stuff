package users

import (
	"context"
	"sukasa/bookings/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByUsernameAndPassword(username, password string) (User, bool)
	FindByID(id primitive.ObjectID) User
}

type userRepository struct {
	client *mongo.Client
}

func (r userRepository) FindByUsernameAndPassword(username, password string) (User, bool) {
	collection := r.client.Database("sukasa").Collection("users")
	filter := bson.M{"username": username, "password": password}

	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return User{}, false
	}

	return result, true
}

func (r userRepository) FindByID(id primitive.ObjectID) User {
	collection := r.client.Database("sukasa").Collection("users")
	filter := bson.M{"_id": id}

	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return User{}
	}

	return result
}

func GetRepository() UserRepository {
	client := db.GetClient()
	return userRepository{client}
}
