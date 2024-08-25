package main

import (
	"Todolist/database"
	"Todolist/logger"
	"Todolist/package/controllers"
)

func main() {
	logger.Init()
	err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	err = database.Migrate()
	if err != nil {
		panic(err)
	}
	err = controllers.Run()
	if err != nil {
		return
	}
}
