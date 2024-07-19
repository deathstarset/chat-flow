package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deathstarset/backend-chatflow/controllers"
	"github.com/deathstarset/backend-chatflow/models"
	"github.com/deathstarset/backend-chatflow/utils"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	user, ok := utils.ParseUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateProfile(c *gin.Context) {
	user, ok := utils.ParseUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse user"})
		return
	}

	var body struct {
		Name     string                `form:"name" binding:"required"`
		Username string                `form:"username" binding:"required"`
		Email    string                `form:"email" binding:"required"`
		Password string                `form:"password" binding:"required"`
		Image    *multipart.FileHeader `form:"image" binding:"required"`
	}

	// delete the old image
	oldImageDst := filepath.Join("uploads", "profile", user.Image)
	err := os.Remove(oldImageDst)
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

	controllers.EditUser(user, models.User{
		ID:       user.ID,
		Name:     body.Name,
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
		Image:    imageName,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated succefully"})
}

func DeleteProfile(c *gin.Context) {
	// getting the session id
	sessionID, err := c.Cookie("session-id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No session ID provided"})
		return
	}
	// getting the user
	user, ok := utils.ParseUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse user"})
		return
	}
	// removing the user from the database
	err = controllers.RemoveUser(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("delete user err : %s", err.Error())})
		return
	}
	// removing the session from redis
	err = utils.RemoveUserSession(sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("remove session err : %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted succefully"})
}
