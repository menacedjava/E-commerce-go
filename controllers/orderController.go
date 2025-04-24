package controllers

import (
	"github.com/menacedjava/helper"
	"github.com/menacedjava/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateOrder Create Order
// @Summary Create Order
// @Description Creates a new order.
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.Order true "Order object"
// @Success 201 {object} object "{"message": "Created Successfully", "id": string}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /order [post]
func CreateOrder(c *gin.Context) {

	order := &models.Order{}
	if err := c.BindJSON(&order); err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, er := order.CreateOrder()
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created Successfully",
		"id":      id,
	})

}

// GetAllOrders Get All Orders
// @Summary Get All Orders
// @Description Retrieves a list of all orders.
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} object "{"message": "All orders", "orders": []}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /order [get]
func GetAllOrders(c *gin.Context) {

	orders, er := models.GetAllOrder()
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all order",
		"orders":  orders,
	})
}

// GetOrder Get Order
// @Summary Get Order
// @Description Retrieves an order by ID.
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} object "{"message": "Success", "order": object}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /order/{id} [get]
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	_id, _ := primitive.ObjectIDFromHex(id)
	order, eror := models.GetAOrder(_id)
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, eror.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"order":   order,
	})
}

// UpdateOrde Update Order
// @Summary Update Order
// @Description Updates the status of an order by ID.
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body models.Order true "Order object"
// @Success 200 {object} object "{"message": "Updated successfully", "count": integer}"
// @Failure 400 {object} object "{"message": "Order Status is empty"}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /order/{id} [put]
func UpdateOrde(c *gin.Context) {
	var order models.Order

	if er := c.BindJSON(&order); er != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, er.Error())
		return
	}
	if order.Status == "" {
		helper.ErrorHandler(c, http.StatusBadRequest, "Order Status is empty")
		return
	}
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)

	count, eror := models.UpdateOrderStatus(id, order.Status)
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, eror.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated successfuly",
		"count":   count,
	})

}
