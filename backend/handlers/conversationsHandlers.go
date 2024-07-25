package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deathstarset/backend-chatflow/controllers"
	"github.com/deathstarset/backend-chatflow/models"
	"github.com/deathstarset/backend-chatflow/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	participants, err := controllers.FindConversationsByParticipants(user.ID, user2.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("find conversations error : %s", err.Error())})
		return
	}
	if participants == 2 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot create conversation between users Twice"})
		return
	}

	conversation := models.Conversation{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		Users:     []models.User{user, user2},
	}

	err = controllers.AddConversation(conversation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("add conversation err : %s", err.Error())})
		return
	}

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

func GetConversation(c *gin.Context) {
	conversationID := c.Param("id")

	user, ok := utils.ParseUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse user"})
		return
	}
	conversations, err := controllers.FindConversationsByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Find conversation error : %s", err.Error())})
		return
	}

	exists := false
	var conversation models.Conversation
	for _, conv := range conversations {
		if conv.ID == conversationID {
			exists = true
			conversation = conv
		}
	}

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Conversation not the user's or conversation not found"})
		return
	}

	conversation, err = controllers.FindConversationByID(conversation.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Find conversation err : %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"conversation": conversation})
}

func DeleteConversation(c *gin.Context) {
	conversationID := c.Param("id")

	user, ok := utils.ParseUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse user"})
		return
	}
	conversations, err := controllers.FindConversationsByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Find conversation error : %s", err.Error())})
		return
	}

	exists := false
	var conversation models.Conversation
	for _, conv := range conversations {
		if conv.ID == conversationID {
			exists = true
			conversation = conv
		}
	}

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Conversation not the user's or conversation not found"})
		return
	}

	err = controllers.RemoveConversation(conversation.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("Delete conversation err : %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversation deleted succefully"})
}
