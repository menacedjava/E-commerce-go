package controllers

import (
	"fmt"
	"github.com/menacedjava/helper"
	"github.com/menacedjava/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUser Register User
// @Summary Create a new User
// @Description Responds with the message and token books as JSON.
// @Tags User
// @Param user body models.User true "User object"
// @Produce json
// @Success 200 {object} object "{"message": "User Created successfully", "token": string}"
// @Failure 400 {object}  object "{"message": "Bad request"} "Bad request"
// @Failure 500 {object} object "{"message": "Internal Server Error"} "Internal Server Error"
// @Router /user/register [post]
func RegisterUser(c *gin.Context) {
	user := &models.User{}
	if eror := c.ShouldBindJSON(&user); eror != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	if user.FName == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	hashPass, hashError := helper.HashPassword(user.Password)
	if hashError != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Error during Hashing passsword")
		return
	}
	user.Password = hashPass
	user.Role = "user"
	id, err := user.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	helper.SendToken(c, id, "user created Successfully")

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "User Created succesfully",
	//	"id":      id,
	//})
}

// LoginUser Login User
// LoginUser Login User
// @Summary Login User
// @Description Responds with the message and token books as JSON.
// @Tags User
// @Param user body models.User true "User object"
// @Produce json
// @Success 200 {object} object "{"message": "Login Successfully", "token": string}"
// @Failure 400 {object}  object "{"message": "Inavalid data"} "Bad request"
// @Failure 401 {object} object "{"message": "No user Found"} "Unauthorized er Error"
// @Router /user/register [post]
func LoginUser(c *gin.Context) {

	var user models.User
	if eror := c.BindJSON(&user); eror != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, "Inavalid data")
		return
	}
	if user.Password == "" || user.Email == "" {
		helper.ErrorHandler(c, http.StatusBadRequest, "please provide email and password")
		return
	}
	foundUser, err := models.UserLogin(user.Email)
	if err != nil {
		helper.ErrorHandler(c, http.StatusUnauthorized, "No user Found ")
		return
	}
	fmt.Println("hhhhhhhh", foundUser)
	compareEror := helper.CompareHashedPassword(user.Password, foundUser.Password)
	if compareEror != nil {
		helper.ErrorHandler(c, http.StatusUnauthorized, "wrong password")
		return
	}
	_id := foundUser.ID.Hex()
	helper.SendToken(c, _id, "Login Successfully")
}

// GetUser Get User By Id
// GetUser By Id
// @Summary Get User By Id
// @Description Retrieves user information by ID.
// @Tags User
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} object "{"message": "User information"}"
// @Failure 404 {object} object "{"message": "No user Found"}"
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {

	id := c.Param("id")

	user, err := models.GetById(id)
	if err != nil {
		helper.ErrorHandler(c, http.StatusNotFound, "No user Found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

// GetAllUsers Get All User
// GetAllUsers Get All Users
// @Summary Get All Users
// @Description Retrieves a list of all users.
// @Tags User
// @Produce json
// @Success 200 {object} object "{"message": "All user list", "total": integer, "users": []}"
// @Failure 404 {object} object "{"message": "No users Found"}"
// @Router /user [get]
func GetAllUsers(c *gin.Context) {

	users, err := models.GetAllUser()
	if err != nil {
		helper.ErrorHandler(c, http.StatusNotFound, "No users Found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "all user list",
		"total":   len(users),
		"users":   users,
	})
}

// UpdateUser UPadte User Profile
// UpdateUser Update User
// @Summary Update User
// @Description Updates user information.
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body models.User true "User object"
// @Success 200 {object} object "{"message": "Updated successfully", "count": integer}"
// @Failure 400 {object} object "{"message": "Invalid info"}"
// @Failure 500 {object} object "{"message": "Error during updating user"}"
// @Router /user [put]
func UpdateUser(c *gin.Context) {

	reqUser, _ := c.Get("user")
	userObj, _ := reqUser.(*models.User)
	var user models.User
	if er := c.BindJSON(&user); er != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, "Inavlid info...")
		return
	}
	//if user.Email == "" || user.Password == "" || user.FName == "" {
	//	helper.ErrorHandler(c, http.StatusBadRequest, "Provide all info")
	//	return
	//}

	id, err := models.UpdateProfile(user, userObj.ID)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Error  during updating user")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated Sucessfuly",
		"count":   id,
	})
}

// LogoutUser Logout User
// @Summary Logout User
// @Description Logs out the user by removing the token cookie.
// @Tags User
// @Produce json
// @Success 200 {object} object "{"message": "Logout successfully"}"
// @Router /user/logout [get]
func LogoutUser(c *gin.Context) {
	c.SetCookie("token", "", 1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout succesfully",
	})
}

// GetProfile GEt Profile
// GetProfile Get User Profile
// @Summary Get User Profile
// @Description Retrieves the profile information of the authenticated user.
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object "{"user": "User profile information"}"
// @Failure 401 {object} object "{"message": "Unauthorized"}"
// @Router /user/profile [get]
func GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		helper.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
