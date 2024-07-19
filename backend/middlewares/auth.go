package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deathstarset/backend-chatflow/controllers"
	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/models"
	"github.com/deathstarset/backend-chatflow/utils"
	"github.com/gin-gonic/gin"
)

func AdminOnly(c *gin.Context) {

	sessionID, err := c.Cookie("session-id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No session ID provided"})
		return
	}

	userSessionString, err := initializers.RD.Get(sessionID).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve session"})
		return
	}

	var userSession utils.SessionUser
	err = json.Unmarshal([]byte(userSessionString), &userSession)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse session"})
		return
	}

	user, err := controllers.FindUserByID(userSession.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to find user"})
		return
	}

	if user.Role != models.Admin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "User is not an admin"})
		return
	}

	c.Next()
}

func AuthOnly(c *gin.Context) {
	log.Println("Auth only middleware")
	sessionID, err := c.Cookie("session-id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No session ID provided"})
		return
	}

	userSessionString, err := initializers.RD.Get(sessionID).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve session"})
		return
	}

	var userSession utils.SessionUser
	err = json.Unmarshal([]byte(userSessionString), &userSession)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse session"})
		return
	}

	user, err := controllers.FindUserByID(userSession.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to find user"})
		return
	}

	c.Set("user", user)

	c.Next()
}
