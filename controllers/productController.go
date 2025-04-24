package controllers

import (
	"fmt"
	"github.com/menacedjava/helper"
	"github.com/menacedjava/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"path/filepath"
	"strconv"
)

// CreateProduct Create Product
// @Summary Create Product
// @Description Creates a new product.
// @Tags Product
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param name formData string true "Product Name"
// @Param description formData string false "Product Description"
// @Param price formData number true "Product Price"
// @Param color formData []string false "Product Colors"
// @Param file formData file true "Product Images"
// @Success 201 {object} object "{"message": "Created Product Successfully", "productId": string, "id": string}"
// @Failure 400 {object} object "{"message": "Bad Request"}"
// @Failure 500 {object} object "{"message": "Error during creation of product"}"
// @Router /product/ [post]
func CreateProduct(c *gin.Context) {
	product := &models.Product{}
	reqUser, _ := c.Get("user")
	user := reqUser.(*models.User)
	name := c.PostForm("name")
	product.Description = c.PostForm("description")
	reqPrice := c.PostForm("price")
	price, _ := strconv.ParseFloat(reqPrice, 64)
	product.Price = price
	product.CreatedBy = user.ID
	product.Color = c.PostFormArray("color")

	files, ok := c.Request.MultipartForm.File["file"]
	if !ok {
		helper.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}
	uploadPath := "./resources/products/images"
	var fileName []string
	for _, file := range files {
		filename := fmt.Sprintf("%s_%s", name, file.Filename)
		filePath := filepath.Join(uploadPath, filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			helper.ErrorHandler(c, http.StatusInternalServerError, "Erro during uploading files")
			return
		}
		fileName = append(fileName, filePath)
	}

	if name != "" {
		product.Name = name
	}
	product.Images = fileName

	id, eror := product.CreateProduct()
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Eror during creation of product")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":   "Created Product Successfully",
		"productId": fileName,
		"id":        id,
	})

}

// GetAllProduct Get All Products
// @Summary Get All Products
// @Description Retrieves a list of all products.
// @Tags Product
// @Accept json
// @Produce json
// @Param lt query string false "Price less than"
// @Param gt query string false "Price greater than"
// @Param color query string false "Color filter"
// @Success 200 {object} object "{"message": "All products", "products": []}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /product/ [get]
func GetAllProduct(c *gin.Context) {
	query := bson.D{}
	lessPrice := c.Query("lt")
	greatePrice := c.Query("gt")
	color := c.Query("color")
	if lessPrice != "" {
		less, _ := strconv.Atoi(lessPrice)
		query = append(query, bson.E{"price", bson.D{
			{"$lt", less},
		}})
	}
	if greatePrice != "" {
		greate, _ := strconv.Atoi(greatePrice)
		query = append(query, bson.E{"price", bson.D{{
			"$gt", greate},
		}})
	}
	if color != "" {
		query = append(query, bson.E{"color", color})
	}
	products, err := models.GetProducts(query)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "all products",
		"products": products,
	})
}

// GetProductById Get Product By ID
// @Summary Get Product By ID
// @Description Retrieves product information by ID.
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} object "{"message": "Get product by id", "product": object}"
// @Failure 404 {object} object "{"message": "No Product Found"}"
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) {
	id := c.Param("id")
	product, er := models.GetByID(id)
	if er != nil {
		helper.ErrorHandler(c, http.StatusNotFound, "No Product Found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get product by id",
		"product": product,
	})
}

// UpdateProduct Update Product
// @Summary Update Product
// @Description Updates product information by ID.
// @Tags Product
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Product ID"
// @Param product body models.Product true "Product object"
// @Success 200 {object} object "{"message": "Updated Successfully", "count": integer}"
// @Failure 400 {object} object "{"message": "Bad Request"}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /product/{id} [put]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if er := c.BindJSON(&product); er != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, er.Error())
		return
	}
	_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	updated, eror := models.UpdateProduct(product, _id)
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, eror.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "updated Successfulyy",
		"count":   updated,
	})
}

// DeleteProduct Delete Product
// @Summary Delete Product
// @Description Deletes a product by ID.
// @Tags Product
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Product ID"
// @Success 200 {object} object "{"message": "Product Deleted", "count": integer}"
// @Failure 500 {object} object "{"message": "Internal Server Error"}"
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)

	deletedCount, eror := models.DeleteProduct(id)
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, eror.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Delteed",
		"count":   deletedCount,
	})
}
