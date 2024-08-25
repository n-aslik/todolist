package controllers

import (
	"Todolist/logger"
	"Todolist/models"
	"Todolist/package/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddUsers(c *gin.Context) {
	var newuser models.User
	err := c.BindJSON(&newuser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = service.CreateUser(newuser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Succesful created"})

}
func EditUsers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var user models.User
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = service.UpdateUser(user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update is successful"})

}
func EditUsersPassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var user models.User
	err = service.UpdateUserPassword(user, id)
	if err != nil {
		logger.Error.Printf("[controllers.EditUsersPassword] invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Edit user`s password  is succesfuly"})
}
func DeleteUsers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUsers] invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	err = service.IsDeletedUser(true, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted is succesfuly"})

}
func PrintUsers(c *gin.Context) {
	logger.Info.Printf("Client with ip: [%s] requested list of users\n", c.ClientIP())
	users, err := service.PrintAllUsers(false, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
	logger.Info.Printf("Client with ip: [%s] got list of users\n", c.ClientIP())
}
func PrintUsersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		logger.Error.Printf("[controllers.PrintUsersByID] invalid user_id path parameter: %s\n", c.Param("id"))
		return
	}
	user, err := service.PrintAllUsersByID(false, false, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
