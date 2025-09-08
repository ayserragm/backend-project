package handlers

import (
	"net/http"

	"github.com/ayserragm/backend-project/internal/db"
	"github.com/ayserragm/backend-project/internal/models"
	"github.com/gin-gonic/gin"
)

// GET /users
func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// POST /users
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	db.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
