package users

import (
	"context"
	"sukasa/bookings/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByUsername(username string) (User, bool)
	FindByUsernameAndPassword(username string, password string) (User, bool)
	FindByID(id primitive.ObjectID) User
}

type userRepository struct {
	client *mongo.Client
}

func (r userRepository) findUser(filter bson.M) (User, bool) {
	collection := r.client.Database("sukasa").Collection("users")

	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return User{}, false
	}

	return result, true
}

func (r userRepository) FindByUsername(username string) (User, bool) {
	filter := bson.M{"username": username}
	return r.findUser(filter)
}

func (r userRepository) FindByUsernameAndPassword(username string, password string) (User, bool) {
	filter := bson.M{"username": username, "password": password}
	return r.findUser(filter)
}

func (r userRepository) FindByID(id primitive.ObjectID) User {
	filter := bson.M{"_id": id}
	user, _ := r.findUser(filter)
	return user
}

func GetRepository() UserRepository {
	client := db.GetClient()
	return userRepository{client}
}
