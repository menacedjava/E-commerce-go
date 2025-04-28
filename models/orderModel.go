package models

import (
	"context"
	"github.com/menacedjava/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductId     primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Price         int64              `json:"price" bson:"price"`
	ShippingPrice int64              `json:"shipping_price"  bson:"shipping_price"`
	TotalPrice    int64              `json:"total_price" bson:"total_price"`
	Status        string             `json:"status" bson:"status"`
	ShippingId    primitive.ObjectID `json:"shipping_id" bson:"shipping_id"`
}

func (order *Order) CreateOrder() (string, error) {

	collect := database.ConnectDB()
	inserted, eror := collect.Order.InsertOne(context.TODO(), &order)
	if eror != nil {
		return "", eror
	}
	id := inserted.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func GetAllOrder() ([]Order, error) {
	var order []Order
	collc := database.ConnectDB()

	cursor, err := collc.Order.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if eror := cursor.All(context.TODO(), &order); eror != nil {
		return nil, eror
	}

	return order, nil
}

func GetAOrder(id primitive.ObjectID) (*Order, error) {
	var order Order
	collec := database.ConnectDB()

	aggregate := bson.A{
		bson.D{{
			"$match", bson.D{
				{
					"_id", id,
				},
			},
		}},

		bson.D{
			{
				"$lookup", bson.D{

					{"from", "ShippingInfo"},
					{"localField", "shipping_id"},
					{"foreignField", "_id"},
					{"as", "shippingAddress"},
				},
			},
		}}

	cursor, er := collec.Order.Aggregate(context.TODO(), aggregate)
	if er != nil {
		return nil, er
	}

	eror := cursor.All(context.TODO(), &order)
	if eror != nil {
		return nil, eror
	}
	return &order, nil
}

func UpdateOrderStatus(id primitive.ObjectID, status string) (int64, error) {

	updated := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}
	coll := database.ConnectDB()
	update, eror := coll.Order.UpdateOne(context.TODO(), bson.M{"_id": id}, updated)
	if eror != nil {
		return 0, eror
	}

	return update.ModifiedCount, nil
}
