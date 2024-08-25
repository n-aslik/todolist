package service

import (
	"Todolist/models"
	"Todolist/package/repository"
	"Todolist/utils"
	"fmt"
)

func CreateUser(user models.User) error {
	_, err := repository.GetUserByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
	}
	user.Password = utils.GenerateHash(user.Password)

	err = repository.CreateUser(user)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateUser(user models.User, id int) error {
	err := repository.EditUser(user.FullName, user.Username, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateUserPassword(user models.User, id int) error {
	err := repository.EditUserPassword(user.Password, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func IsDeletedUser(isdeleted bool, id int) error {
	err := repository.DeleteUser(isdeleted, id)
	if err != nil {
		fmt.Println(err)

	}
	return err
}

func PrintAllUsers(isdeleted bool, isblocked bool) (user []models.User, err error) {
	users, err := repository.GetAllUsers(isdeleted, isblocked)
	if err != nil {
		return users, err

	}
	return users, nil
}

func PrintAllUsersByID(isdeleted bool, isblocked bool, id int) (user []models.User, err error) {
	users, err := repository.GetAllUserByID(isdeleted, isblocked, id)
	if err != nil {
		return users, err

	}
	return users, nil
}
