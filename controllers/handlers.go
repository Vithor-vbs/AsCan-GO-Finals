package controllers

import (
	"net/http"

	"github.com/Vithor-vbs/API-Requests-Example/initializers"
	"github.com/Vithor-vbs/API-Requests-Example/models"
	"github.com/gin-gonic/gin"
)

// get ALl
// GetUsers godoc
// @Summary List existing users
// @Description Get all the existing users
// @Accept json
// @Produce json
// @Success 200 {object} []models.User
// @Failure 401 {string} string "Unauthorized"
// @Router /user [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Get user by query parameter
func FindUser(c *gin.Context) {
	var user models.User

	if err := initializers.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// schema that can validate the user’s input to prevent from getting invalid data:
type CreateUserInput struct {
	Username   string `json:"username" binding:"required"`
	FirstName  string `json:"firsName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	UserStatus int    `json:"userStatus"`
}

// POST Method:
// @title Create User
// @Summary Creates a account
// @produce json
// @Description creates an User
// @Param request body CreateUserInput true "User"
// @success 200 {object} models.User
// @failure 400 {object} string
// @Router /user [post]
func CreateUser(c *gin.Context) {

	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create newUser
	user := models.User{Username: input.Username, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, Password: input.Password, Phone: input.Phone, UserStatus: input.UserStatus}
	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST user with array -> createWithArray endpopint
// @title Create User by array
// @Summary Creates an array of users
// @produce json
// @Description creates an array of Users
// @Param request body []CreateUserInput true "User"
// @success 200 {object} []models.User
// @failure 400 {object} string
// @Router /user/createWithArray [post]
func CreateUserByArr(c *gin.Context) {
	var input []CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, v := range input {
		user := models.User{Username: v.Username, FirstName: v.FirstName, LastName: v.LastName, Email: v.Email, Password: v.Password, Phone: v.Phone, UserStatus: v.UserStatus}
		initializers.DB.Create(&user)
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

type UpdateUserInput struct {
	Username   string `json:"username"`
	FirstName  string `json:"firsName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userStatus"`
}

// @title Update User
// @Summary updates a User
// @produce json
// @Description Updates an User
// @Param request body UpdateUserInput true "User"
// @success 200 {object} models.User
// @failure 400 {object} string
// @Router /user/:id [patch]
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @title Update User
// @Summary updates a User
// @produce json
// @Description Updates an User
// @Param request body UpdateUserInput true "User"
// @success 200 {object} models.User
// @failure 400 {object} string
// @Router /user/:id [put]
func PutUser(c *gin.Context) {
	id := c.Param("id")
	var input UpdateUserInput

	c.Bind(&input)

	var post models.User
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.User{
		Username:   input.Username,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Email:      input.Email,
		Password:   input.Password,
		Phone:      input.Phone,
		UserStatus: input.UserStatus,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// Delete a User
// @title Delete User
// @Summary Eliminates an User
// @Description Delete an User
// @Accept json
// @Produce json
// @Success 200 {object} []models.User
// @Failure 401 {string} string "Unauthorized"
// @Router /user/:id [delete]
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
