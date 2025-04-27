package database

import (
	"context"
	"fmt"
	"github.com/menacedjava/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

type Collections struct {
	UserCollection    *mongo.Collection
	ProductCollection *mongo.Collection
	ReviewCollection  *mongo.Collection
	Order             *mongo.Collection
	ShipingAddres     *mongo.Collection
}

var (
	once       sync.Once
	collection *Collections
)

//Create Connection to datbase

func ConnectDB() *Collections {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("env file load error")
		}
		uri := os.Getenv("MONGO_URL")
		if uri == "" {
			log.Fatal("Mongo Url is empty")
		}

		connection, eror := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if eror != nil {
			log.Fatal("Connection to mongo db failed")
		}

		db := connection.Database("shopping")
		userColl := db.Collection("User")
		productColl := db.Collection("Product")
		reviewColl := db.Collection("Review")
		orderColl := db.Collection("Order")
		shippingColl := db.Collection("ShippingInfo")
		collection = &Collections{
			UserCollection:    userColl,
			ProductCollection: productColl,
			ReviewCollection:  reviewColl,
			Order:             orderColl,
			ShipingAddres:     shippingColl,
		}

		fmt.Println("Connection to Db Successfull..........")
	})
	return collection

}
