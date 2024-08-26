package main

import (
	"Todolist/configs"
	"Todolist/database"
	"Todolist/logger"
	"Todolist/package/controllers"
	"errors"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
	}
	err = configs.ReadString()
	if err != nil {
		panic(err)
	}
	logger.Init()
	err = database.ConnectDB()
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
