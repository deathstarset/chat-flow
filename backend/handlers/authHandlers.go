package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/deathstarset/backend-chatflow/controllers"
	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/models"
	"github.com/deathstarset/backend-chatflow/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("body validation err : %s", err.Error())})
		return
	}

	user, err := controllers.FindUserByUsername(body.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user find err : %s", err.Error())})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("wrong password err : %s", err.Error())})
		return
	}

	err = utils.CreateUserSession(initializers.RD, c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("session err : %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Register(c *gin.Context) {
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

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Password hash err : %s", err.Error())})
		return
	}
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
