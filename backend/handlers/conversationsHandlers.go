package handlers

import (
	"fmt"
	"net/http"

	"github.com/deathstarset/backend-chatflow/controllers"
	"github.com/deathstarset/backend-chatflow/models"
	"github.com/deathstarset/backend-chatflow/utils"
	"github.com/gin-gonic/gin"
)

func CreateConversation(c *gin.Context) {
	user, ok := utils.ParseUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse user"})
		return
	}
	var body struct {
		UserID string `json:"user_id" binding:"required"`
	}
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Validation err : %s", err.Error())})
		return
	}

	user2, err := controllers.FindUserByID(body.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("User not found : %s", err.Error())})
		return
	}

	controllers.FindConversationsByParticipants(user.ID, user2.ID)

	/* conversation := models.Conversation{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		Users:     []models.User{user, user2},
	}

	err = controllers.AddConversation(conversation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("add conversation err : %s", err.Error())})
		return
	} */

	c.JSON(http.StatusCreated, gin.H{"message": "conversation created succefully"})
}

func GetAllConversations(c *gin.Context) {
	var conversations []models.Conversation
	conversations, err := controllers.FindAllConversations()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("find conversations err : %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"conversations": conversations})
}
