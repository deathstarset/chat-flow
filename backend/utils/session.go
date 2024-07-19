package utils

import (
	"encoding/json"
	"time"

	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/models"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type SessionUser struct {
	ID string `json:"id"`
}

func CreateUserSession(c *gin.Context, user models.User) error {
	sessionID := uuid.NewString()
	sessionUser := SessionUser{
		ID: user.ID,
	}

	// turn the struct into json
	jsonSession, err := json.Marshal(sessionUser)
	if err != nil {
		return err
	}

	// add the session to redis
	_, err = initializers.RD.Set(sessionID, jsonSession, time.Hour*1).Result()
	if err != nil {
		return err
	}

	// send a cookie with the session id
	c.SetCookie("session-id", sessionID, 180, "/", "localhost", false, true)
	return nil
}

func ParseUser(c *gin.Context) (models.User, bool) {
	var user models.User
	userContext, ok := c.Get("user")
	if !ok {
		return user, ok
	}
	user, ok = userContext.(models.User)
	if !ok {
		return user, ok
	}
	return user, true
}

func RemoveUserSession(sessionID string) error {
	_, err := initializers.RD.Del(sessionID).Result()
	if err != nil {
		return err
	}
	return nil
}
