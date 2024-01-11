package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sahildotexe/go-gin-gorm-rest/config"
	"github.com/sahildotexe/go-gin-gorm-rest/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(201, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Where("id = ?", id).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Where("id = ?", id).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Where("id = ?", id).First(&user)
	c.JSON(200, &user)
}
