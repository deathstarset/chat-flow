package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/deathstarset/backend-chatflow/controllers"
	"github.com/deathstarset/backend-chatflow/models"
	"github.com/deathstarset/backend-chatflow/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {
	var body struct {
		Name     string                `form:"name" binding:"required"`
		Username string                `form:"username" binding:"required"`
		Email    string                `form:"email" binding:"required"`
		Password string                `form:"password" binding:"required"`
		Image    *multipart.FileHeader `form:"image" binding:"required"`
	}
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("body validation err: %s", err.Error())})
		return
	}

	image := body.Image
	imageName, err := utils.SaveFile(image, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("File upload err : %s", err.Error())})
		return
	}

	// hashing the password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Password hash err : %s", err.Error())})
		return
	}

	// saving data into the database
	err = controllers.AddUser(models.User{
		ID:        uuid.NewString(),
		Name:      body.Name,
		Username:  body.Username,
		Email:     body.Email,
		Password:  hashedPassword,
		Image:     imageName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Db err : %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created succefully"})
}

func GetUser(c *gin.Context) {
	userId := c.Param("id")
	user, err := controllers.FindUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("find user err : %s", err.Error())})
		return
	}
	c.JSON(200, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	var user models.User
	var body struct {
		Name     string                `form:"name" binding:"required"`
		Username string                `form:"username" binding:"required"`
		Email    string                `form:"email" binding:"required"`
		Password string                `form:"password" binding:"required"`
		Image    *multipart.FileHeader `form:"image" binding:"required"`
	}
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("body validation err: %s", err.Error())})
		return
	}

	// getting the user that we are gonna update by id
	user, err = controllers.FindUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// delete the old image
	oldImageDst := filepath.Join("uploads", "profile", user.Image)
	err = os.Remove(oldImageDst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Server err : %s", err.Error())})
		return
	}

	// upload the image
	image := body.Image
	imageName, err := utils.SaveFile(image, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("File upload err : %s", err.Error())})
		return
	}

	// hash the password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Password hash err : %s", err.Error())})
		return
	}

	// update the user
	controllers.EditUser(user, models.User{
		ID:       user.ID,
		Name:     body.Name,
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
		Image:    imageName,
	})

	c.JSON(200, gin.H{"message": "User updated succefully"})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	err := controllers.RemoveUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("delete user err : %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted succefully"})
}

func GetAllUsers(c *gin.Context) {
	users, err := controllers.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("get users err : %s", err.Error())})
		return
	}
	c.JSON(200, gin.H{"users": users})
}
