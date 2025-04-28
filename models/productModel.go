package models

import (
	"context"
	"github.com/menacedjava/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	Price         float64            `json:"price,omitempty" bson:"price,omitempty"`
	Color         []string           `json:"color" bson:"color"`
	Images        []string           `json:"images" bson:"images"`
	CreatedBy     primitive.ObjectID `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	AverageReview float64            `json:"averageReview" bson:"averageReview"`
}

func (product *Product) CreateProduct() (string, error) {
	productColl := database.ConnectDB()

	inserted, err := productColl.ProductCollection.InsertOne(context.TODO(), &product)
	if err != nil {
		return "", err
	}
	_id := inserted.InsertedID.(primitive.ObjectID)
	id := _id.Hex()

	return id, nil
}

func GetProducts(query bson.D) ([]Product, error) {
	var product []Product
	productColl := database.ConnectDB()
	var (
		cursor *mongo.Cursor
		er     error
	)
	if query != nil {
		cursor, er = productColl.ProductCollection.Find(context.TODO(), query)
	} else {
		cursor, er = productColl.ProductCollection.Find(context.TODO(), bson.M{})
	}
	if er != nil {
		return nil, er

	}
	if eror := cursor.All(context.TODO(), &product); eror != nil {
		return nil, eror
	}

	return product, nil
}

func GetByID(id string) (*Product, error) {
	productColl := database.ConnectDB()
	var product Product
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	err := productColl.ProductCollection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil

}

func DeleteProduct(id primitive.ObjectID) (int64, error) {

	collect := database.ConnectDB()

	deletedCount, err := collect.ProductCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return 0, err
	}

	_, eror := collect.ReviewCollection.DeleteMany(context.TODO(), bson.M{"productId": id})
	if eror != nil {
		return 0, eror
	}

	return deletedCount.DeletedCount, nil
}

func UpdateProduct(product Product, id primitive.ObjectID) (int64, error) {

	updated := bson.D{}

	if product.Name != "" {
		updated = append(updated, bson.E{"name", product.Name})
	}
	if product.Description != "" {
		updated = append(updated, bson.E{"description", product.Description})
	}
	if product.Price != 0 {
		updated = append(updated, bson.E{"price", product.Price})
	}
	if len(product.Color) != 0 {
		updated = append(updated, bson.E{"color", product.Color})
	}

	collec := database.ConnectDB()
	updatedOne, err := collec.ProductCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, updated)
	if err != nil {
		return 0, err
	}

	return updatedOne.ModifiedCount, nil
}
