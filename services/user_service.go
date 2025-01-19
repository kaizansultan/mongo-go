package services

import (
	"context"
	"project/database"
	"project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	collection := database.GetCollection("users")
	result, err := collection.InsertOne(context.Background(), user)
	return result, err
}

func GetUserByEmail(email string) (models.User, error) {
	collection := database.GetCollection("users")
	var user models.User
	filter := bson.M{"email": email}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}
