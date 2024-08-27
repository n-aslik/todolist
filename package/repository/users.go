package repository

import (
	"Todolist/database"
	"Todolist/logger"
	"Todolist/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	err := database.GetconnectDB().Create(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.createuser]error in added user %s\n", err.Error())
	}
	return nil
}
func EditUser(fullname, username string, id int) error {
	err := database.GetconnectDB().Omit("password").Where("id=?", id).Updates(models.User{FullName: fullname, Username: username}).Error
	if err != nil {
		logger.Error.Printf("[repository.edituser]error in updated user %s\n", err.Error())
	}
	return nil
}

func DeleteUser(isdeleted bool, id int) error {
	err := database.GetconnectDB().Model(&models.User{}).Where("id=?", id).Update("is_deleted", isdeleted).Error
	if err != nil {
		logger.Error.Printf("[repository.deleteuser]error in deleted user %s\n", err.Error())
	}
	return nil
}
func EditUserPassword(password string, id int) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error.Printf("[repository.edituserpassword]error in hashing password: %s\n", err.Error())
		return err
	}
	err = database.GetconnectDB().Model(&models.User{}).Where("id=?", id).Update("password", passwordHash).Error
	if err != nil {
		logger.Error.Printf("[repository.edituserpassword]error in updated user password %s\n", err.Error())
	}
	return nil
}

func GetAllUsers(isdeleted bool, isblocked bool) (user []models.User, err error) {
	err = database.GetconnectDB().Where("is_deleted=?", isdeleted).Where("is_blocked=?", isblocked).Find(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.getallusers]error in getting all users %s\n", err.Error())
		return user, err
	}
	return user, nil
}
func GetAllUserByID(isdeleted bool, isblocked bool, id int) (user []models.User, err error) {
	err = database.GetconnectDB().Where("is_deleted=?", isdeleted).Where("is_blocked=?", isblocked).Where("id=?", id).Find(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.getallusersbyid]error in getting all users by id %s\n", err.Error())
		return user, err
	}
	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = database.GetconnectDB().Where("username=?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[service.getuserbyusername]error in getting user by username  %s\n", err.Error())
		return user, err
	}
	return user, nil
}

func GetUserByUsernameAndPassword(username, password string) (user models.User, err error) {
	err = database.GetconnectDB().Where("username=? AND password=?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[service.getuserbyusername]error in getting user by username  %s\n", err.Error())
		return user, err
	}
	return user, nil
}
