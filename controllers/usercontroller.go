package controllers

import (
	"net/http"

	"apigolang/database"
	"apigolang/models"

	"github.com/gin-gonic/gin"
)

type UpdateUserInput struct{
  	Username          string `json:"username" binding:"required"`
	FirstName         string `json:"first_name" binding:"required"`
	LastName          string `json:"last_name" binding:"required"`
	MobilePhoneNumber string `json:"mobile_phone_number" binding:"required"`
	Email             string `json:"email" gorm:"unique" binding:"required"`
	Birthday          string `json:"date"`
}

func GetAllUsers(c *gin.Context){
  var users []models.User
  database.InstanceData.Find(&users)
  c.JSON(http.StatusOK, gin.H{"data":users})
}

func CreateUser(c *gin.Context) {
  // Validate input
  var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.InstanceData.Create(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

func GetUserByID(c *gin.Context) {  // Get model if exist
  var user models.User
  id := c.Request.URL.Query().Get("id")
  if err := database.InstanceData.Where("id = ?", id).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})

}

func UpdateUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  id := c.Request.URL.Query().Get("id")
  if err := database.InstanceData.Where("id = ?", id).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input UpdateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  database.InstanceData.Model(&user).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context){
   var user models.User

	if err := database.InstanceData.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
        return
    }

    database.InstanceData.Delete(&user)

    c.JSON(http.StatusOK, gin.H{"data": true})
}

