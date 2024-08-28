package controllers

import (
	"Todolist/logger"
	"Todolist/models"
	"Todolist/package/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	var newtask models.Task
	err := c.BindJSON(&newtask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = service.AddTask(newtask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Succesful created"})

}

func GetAllTasks(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	isDoneStr := c.Query("is_done")
	isDone, err := strconv.ParseBool(isDoneStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	tasks, err := service.PrintAllTasks(false, isDone, false, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetAllTasksByID(c *gin.Context) {
	uid := c.GetUint(userIDCtx)
	if uid == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	tid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetTasksByID] invalid task_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	task, err := service.PrintAllTasksByID(false, false, uint(tid), uint(uid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func UpdateTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.UpdateTaskByID] invalid task_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var task models.Task
	err = c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = service.UpdateTask(task, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update is successful"})

}

func ChecksasDone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.CheckasDone] invalid task_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	err = service.CheckTaskasDone(true, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check as done is succesfuly"})

}
func DeleteTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTaskByID] invalid task_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	err = service.DeleteTask(true, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted is succesfuly"})

}
