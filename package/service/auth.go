package service

import (
	"Todolist/package/repository"
	"Todolist/utils"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}

	accessToken, err = GenerateToken(uint(user.ID), user.Username)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
