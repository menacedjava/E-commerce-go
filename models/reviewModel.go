package models

import (
	"context"
	"fmt"
	"github.com/menacedjava/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
)

type Review struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description"`
	Rating      float64            `json:"rating"`
	ProductId   primitive.ObjectID `json:"productId" bson:"productId"`
}

func (review *Review) CreateReview() (string, error) {
	reviewColl := database.ConnectDB()

	inserted, er := reviewColl.ReviewCollection.InsertOne(context.TODO(), &review)
	if er != nil {
		return "", er
	}
	pipeLine := []bson.M{
		{
			"$match": bson.M{"productId": review.ProductId},
		},
		{
			"$group": bson.M{

				"_id": "$productId",
				"avgRating": bson.M{
					"$avg": "$rating",
				},
			},
		},
	}

	cursor, err := reviewColl.ReviewCollection.Aggregate(context.TODO(), pipeLine)
	if err != nil {
		return "", err
	}

	// Decode and print the results
	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		return "", err
	}

	// Print the average rating results
	for _, result := range results {
		productId := result["_id"]
		avgRating := result["avgRating"].(float64)
		roundedAvgRating := math.Round(avgRating*100) / 100 // Round to two decimal places
		fmt.Printf("ProductId: %s, Average Rating: %.2f\n", productId, roundedAvgRating)
		filter := bson.M{
			"_id": productId,
		}
		updated := bson.M{
			"$set": bson.M{
				"averageReview": roundedAvgRating,
			},
		}
		_, err := reviewColl.ProductCollection.UpdateOne(context.TODO(), filter, updated)
		if err != nil {
			return "", err
		}

	}
	_id := inserted.InsertedID.(primitive.ObjectID)
	id := _id.Hex()
	return id, nil
}

func Delete(id primitive.ObjectID) (int64, error) {
	collec := database.ConnectDB()
	var review Review
	filter := bson.M{"_id": id}

	eror := collec.ReviewCollection.FindOne(context.TODO(), filter).Decode(&review)
	if eror != nil {
		return 0, eror
	}
	delted, er := collec.ReviewCollection.DeleteOne(context.TODO(), filter)
	if er != nil {
		return 0, er
	}
	pipeLine := []bson.M{
		{
			"$match": bson.M{"productId": review.ProductId},
		},
		{
			"$group": bson.M{

				"_id": "$productId",
				"avgRating": bson.M{
					"$avg": "$rating",
				},
			},
		},
	}

	cursor, err := collec.ReviewCollection.Aggregate(context.TODO(), pipeLine)
	if err != nil {
		return 0, err
	}

	// Decode and print the results
	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		return 0, err
	}

	// Print the average rating results
	for _, result := range results {
		productId := result["_id"]
		avgRating := result["avgRating"].(float64)
		roundedAvgRating := math.Round(avgRating*100) / 100 // Round to two decimal places
		fmt.Printf("ProductId: %s, Average Rating: %.2f\n", productId, roundedAvgRating)
		filter := bson.M{
			"_id": productId,
		}
		updated := bson.M{
			"$set": bson.M{
				"averageReview": roundedAvgRating,
			},
		}
		_, err := collec.ProductCollection.UpdateOne(context.TODO(), filter, updated)
		if err != nil {
			return 0, err
		}

		return 0, nil
	}
	return delted.DeletedCount, nil
}

func GetReviews() ([]Review, error) {
	coll := database.ConnectDB()
	var reviews []Review
	cursor, er := coll.ReviewCollection.Find(context.TODO(), bson.M{})
	if er != nil {
		return nil, er
	}
	eror := cursor.All(context.TODO(), &reviews)
	if eror != nil {
		return nil, eror
	}
	return reviews, nil

}

func UpdateRev(review Review, id primitive.ObjectID) (int64, error) {

	coll := database.ConnectDB()
	filter := bson.M{
		"_id": id,
	}
	updateFields := bson.D{}
	if review.Title != "" {

		updateFields = append(updateFields, bson.E{"title", review.Title})

	}
	if review.Description != "" {
		updateFields = append(updateFields, bson.E{"description", review.Description})
	}
	if review.Rating != 0 {
		updateFields = append(updateFields, bson.E{"rating", review.Rating})
	}

	upadtedCount, er := coll.ReviewCollection.UpdateOne(context.TODO(), filter, updateFields)
	if er != nil {
		return 0, nil
	}
	if review.Rating != 0 {
		var rev Review
		eror := coll.ReviewCollection.FindOne(context.TODO(), filter).Decode(&rev)
		if eror != nil {
			return 0, eror
		}
		pipeLine := []bson.M{
			{
				"$match": bson.M{"productId": rev.ProductId},
			},
			{
				"$group": bson.M{

					"_id": "$productId",
					"avgRating": bson.M{
						"$avg": "$rating",
					},
				},
			},
		}

		cursor, err := coll.ReviewCollection.Aggregate(context.TODO(), pipeLine)
		if err != nil {
			return 0, err
		}

		// Decode and print the results
		var results []bson.M
		if err := cursor.All(context.TODO(), &results); err != nil {
			return 0, err
		}
		for _, result := range results {
			productId := result["_id"]
			avgRating := result["avgRating"].(float64)
			roundedAvgRating := math.Round(avgRating*100) / 100 // Round to two decimal places
			fmt.Printf("ProductId: %s, Average Rating: %.2f\n", productId, roundedAvgRating)
			filter := bson.M{
				"_id": productId,
			}
			updated := bson.M{
				"$set": bson.M{
					"averageReview": roundedAvgRating,
				},
			}
			_, err := coll.ProductCollection.UpdateOne(context.TODO(), filter, updated)
			if err != nil {
				return 0, err
			}

			return 0, nil
		}
	}

	return upadtedCount.ModifiedCount, nil
}

func GetRev(id primitive.ObjectID) (*Review, error) {
	var review Review
	coll := database.ConnectDB()
	err := coll.ReviewCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&review)
	if err != nil {
		return nil, err
	}

	return &review, nil
}
