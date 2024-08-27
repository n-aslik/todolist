package controllers

import (
	"Todolist/configs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() error {

	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)
	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}
	usersG := router.Group("/users")
	{
		usersG.POST("", AddUsers)
		usersG.GET("", PrintUsers)
		usersG.GET("/:id", PrintUsersByID)
		usersG.PUT("/:id", EditUsers)
		usersG.PATCH("/:id", EditUsersPassword)
		usersG.DELETE("/:id", DeleteUsers)
	}

	tasksG := router.Group("/tasks", checkUserAuthentication)
	{
		tasksG.POST("", AddTask)
		tasksG.GET("", GetAllTasks)
		tasksG.GET("/:id", GetAllTasksByID)
		tasksG.PUT("/:id", UpdateTaskByID)
		tasksG.PATCH("/:id", ChecksasDone)
		tasksG.DELETE("/:id", DeleteTaskByID)
	}

	err := router.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))
	if err != nil {
		panic(err)
	}
	return nil
}
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
