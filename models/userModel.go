package models

import (
	"context"
	"github.com/menacedjava/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FName    string             `json:"f_name"`
	LName    string             `json:"l_name"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
	Gender   string             `json:"gender"`
	PhoneNo  string             `json:"phone_no"`
	Role     string             `json:"role"`
}

func (user *User) CreateUser() (string, error) {

	collection := database.ConnectDB()
	insertId, eror := collection.UserCollection.InsertOne(context.TODO(), &user)

	if eror != nil {
		return "", eror
	}
	_id := insertId.InsertedID.(primitive.ObjectID)
	id := _id.Hex()
	return id, nil
}

func UserLogin(email string) (*User, error) {
	var user User
	collection := database.ConnectDB()

	filter := bson.M{"email": email}

	err := collection.UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func GetById(id string) (*User, error) {
	collection := database.ConnectDB()
	_id, _ := primitive.ObjectIDFromHex(id)
	var user User
	filter := bson.M{"_id": _id}

	err := collection.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUser() ([]User, error) {
	collection := database.ConnectDB()
	var users []User

	cursor, er := collection.UserCollection.Find(context.TODO(), bson.M{})

	if er != nil {
		return nil, er
	}
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateProfile(user User, userId primitive.ObjectID) (int64, error) {
	collection := database.ConnectDB()
	filter := bson.M{"_id": userId}
	updateFields := bson.D{}

	if user.FName != "" {
		updateFields = append(updateFields, bson.E{"fname", user.FName})
	}

	if user.LName != "" {
		updateFields = append(updateFields, bson.E{"lname", user.LName})
	}

	if user.Gender != "" {
		updateFields = append(updateFields, bson.E{"gender", user.Gender})
	}

	if user.PhoneNo != "" {
		updateFields = append(updateFields, bson.E{"phoneno", user.PhoneNo})
	}

	update := bson.M{
		"$set": updateFields,
	}
	updated, err := collection.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}

	return updated.ModifiedCount, nil

}

