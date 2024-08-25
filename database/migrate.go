package database

import (
	"Todolist/models"
	"errors"
)

func Migrate() error {
	err := conn.AutoMigrate(models.Task{}, models.User{})
	if err != nil {
		return errors.New("Failed to begin transaction: " + err.Error())
	}

	return nil
}
