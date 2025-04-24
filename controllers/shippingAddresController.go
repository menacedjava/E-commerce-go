package controllers

import (
	"github.com/menacedjava/helper"
	"github.com/menacedjava/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateShiipingAddress  Create Shipping Address
// @Summary Create Shipping Address
// @Description Creates a new shipping address.
// @Tags Shipping
// @Accept json
// @Produce json
// @Param shipping body models.Shipping true "Shipping object"
// @Success 200 {object} object "{"message": "Created Successfully", "_id": string}"
// @Failure 400 {object} object "{"message": "Bad Request"}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /order/shipping [post]
func CreateShiipingAddress(c *gin.Context) {

	shiiping := &models.Shipping{}

	if er := c.BindJSON(&shiiping); er != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, er.Error())
		return
	}

	id, err := shiiping.CreateShipingAddress()

	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	_id, _ := primitive.ObjectIDFromHex(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Created Succesfuly",
		"_id":     _id,
	})
}
