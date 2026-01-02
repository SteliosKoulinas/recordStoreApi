package handlers

import (
	"github.com/SteliosKoulinas/recordStoreApi/services"

	"github.com/gin-gonic/gin"
)

var UserService = services.UserService{}

func GetUsers(c *gin.Context) {
	users, err := UserService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

/* create User
func CreateUser(c *gin.Context) {
	var input models.Users

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := UserService.Create(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, input)
}*/
