package controllers

import (
	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/models"
)

func FindUserByID(id string) (models.User, error) {
	var user models.User
	result := initializers.DB.Find(&user, "id = ?", id)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func RemoveUser(id string) error {
	result := initializers.DB.Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllUsers() ([]models.User, error) {
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func AddUser(user models.User) error {
	result := initializers.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func EditUser(oldUser models.User, newUser models.User) {

	oldUser.Name = newUser.Name
	oldUser.Username = newUser.Username
	oldUser.Email = newUser.Email
	oldUser.Password = newUser.Password
	oldUser.Image = newUser.Image

	initializers.DB.Save(&oldUser)

}
