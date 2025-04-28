package models

import (
	"context"
	"github.com/menacedjava/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shipping struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FName   string             `json:"f_name" bson:"f_name"`
	LName   string             `json:"l_name" bson:"l_name"`
	Email   string             `json:"email" bson:"email"`
	PhoneNo string             `json:"phone_no"`
	Address string             `json:"address" bson:"address"`
	State   string             `json:"state" bson:"state"`
}

func (shipping *Shipping) CreateShipingAddress() (string, error) {

	coll := database.ConnectDB()
	inserted, err := coll.ShipingAddres.InsertOne(context.TODO(), &shipping)

	if err != nil {
		return "", err
	}

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil
}

