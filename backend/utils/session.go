package utils

import (
	"encoding/json"
	"log"
	"time"

	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/google/uuid"
)

type SessionUser struct {
	ID string `json:"id"`
}

func CreateUserSession(client *redis.Client, c *gin.Context, user models.User) error {
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
	_, err = client.Set(sessionID, jsonSession, time.Hour*1).Result()
	if err != nil {
		return err
	}

	// send a cookie with the session id
	c.SetCookie("session-id", sessionID, 180, "/", "localhost", false, true)
	return nil
}

func AuthMiddleware(c *gin.Context) {
	sessionID, err := c.Cookie("session-id")
	if err != nil {

	}
	userSessionString, err := initializers.RD.Get(sessionID).Result()
	if err != nil {

	}
	var userSession SessionUser
	err = json.Unmarshal([]byte(userSessionString), &userSession)
	if err != nil {

	}
	log.Println(userSession.ID)
	c.Next()
}
