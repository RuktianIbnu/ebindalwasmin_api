package auth

import (
	"errors"
	"../models"
	"../utils"
	"../security"
)

var (
	ErrUserNotFound = errors.New("User Tidak Ditemukan")
)

func SignIn(email, password string) (string, error) {
	user := models.GetUserByEmail(email)
	if user.Id == 0 {
		return "", ErrUserNotFound
	}
	err := security.VerifyPassword(user.Password, password)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}