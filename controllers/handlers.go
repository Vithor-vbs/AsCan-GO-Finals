package controllers

import (
	"net/http"

	"github.com/Vithor-vbs/API-Requests-Example/initializers"
	"github.com/Vithor-vbs/API-Requests-Example/models"
	"github.com/gin-gonic/gin"
)

// get ALl
func GetUsers(c *gin.Context) {
	
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
		
}

// Get user by query parameter
func FindUser(c *gin.Context) {  // Get model if exist
  var user models.User

  if err := initializers.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}

//schema that can validate the user’s input to prevent from getting invalid data:
type CreateUserInput struct {
	Username   string `json:"username" binding:"required"`
	FirstName  string `json:"firsName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	UserStatus int    `json:"userStatus"`
}
  //POST Method:
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

type UpdateUserInput struct{
	Username   string `json:"username"`
	FirstName  string `json:"firsName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userStatus"`
}

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

func PutUser(c *gin.Context){
	id := c.Param("id")

	var UpdateUserInput struct{
		Username   string `json:"username"`
		FirstName  string `json:"firsName"`
		LastName   string `json:"lastName"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Phone      string `json:"phone"`
		UserStatus int    `json:"userStatus"`
	}

	c.Bind(&UpdateUserInput)

	var post models.User
	initializers.DB.First(&post, id)


	initializers.DB.Model(&post).Updates(models.User{
		Username: UpdateUserInput.Username,
		FirstName: UpdateUserInput.FirstName,
		LastName: UpdateUserInput.LastName,
		Email: UpdateUserInput.Email,
		Password: UpdateUserInput.Password,
		Phone: UpdateUserInput.Phone,
		UserStatus: UpdateUserInput.UserStatus,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// Delete a User
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