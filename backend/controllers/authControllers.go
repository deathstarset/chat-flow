package controllers

import (
	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/models"
)

func FindUserByUsername(username string) (models.User, error) {
	var user models.User
	result := initializers.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
