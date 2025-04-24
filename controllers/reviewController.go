package controllers

import (
	"github.com/menacedjava/helper"
	"github.com/menacedjava/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateReview Create Review
// @Summary Create Review
// @Description Creates a new review for a product by ID.
// @Tags Review
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param review body models.Review true "Review object"
// @Success 200 {object} object "{"message": "Review added successfully", "id": string}"
// @Failure 400 {object} object "{"message": "Invalid data"}"
// @Failure 500 {object} object "{"message": "Error during creating review"}"
// @Router /review/{id} [post]
func CreateReview(c *gin.Context) {

	review := &models.Review{}
	id := c.Param("id")
	productId, _ := primitive.ObjectIDFromHex(id)
	if err := c.BindJSON(&review); err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Invalid data")
		return
	}
	review.ProductId = productId

	id, er := review.CreateReview()
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Error during creating review")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "review added succesffuly",
		"id":      id,
	})
}

// DeleteReview Delete Review
// @Summary Delete Review
// @Description Deletes a review by ID.
// @Tags Review
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} object "{"count": integer}"
// @Failure 400 {object} object "{"message": "Please Provide ID"}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /review/{id} [delete]
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ErrorHandler(c, http.StatusBadRequest, "Plaese Provide Id")
		return
	}
	_id, _ := primitive.ObjectIDFromHex(id)

	deletedCount, err := models.Delete(_id)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"count": deletedCount,
	})
}

// GetAllReviews Get All Reviews
// @Summary Get All Reviews
// @Description Retrieves a list of all reviews.
// @Tags Review
// @Accept json
// @Produce json
// @Success 200 {object} object "{"total": integer, "reviews": []}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /review [get]
func GetAllReviews(c *gin.Context) {

	reviews, er := models.GetReviews()
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total":   len(reviews),
		"reviews": reviews,
	})
}

// UpdateReview Update Review
// @Summary Update Review
// @Description Updates a review by ID.
// @Tags Review
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Param review body models.Review true "Review object"
// @Success 200 {object} object "{"updatedCount": integer}"
// @Failure 400 {object} object "{"message": "Provide the valid data"}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /review/{id} [put]
func UpdateReview(c *gin.Context) {

	var review models.Review
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)
	if err := c.BindJSON(&review); err != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, "Provide the valid data")
		return
	}

	updated, er := models.UpdateRev(review, id)
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"updatedCount": updated,
	})

}

// GetRevById Get Review By ID
// @Summary Get Review By ID
// @Description Retrieves a review by ID.
// @Tags Review
// @Accept json
// @Produce json
// @Param id path string true "Review ID"
// @Success 200 {object} object "{"message": "Got successfully", "review": object}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /review/{id} [get]
func GetRevById(c *gin.Context) {
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)

	review, er := models.GetRev(id)
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Got succesfuly",
		"review":  review,
	})
}
