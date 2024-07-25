package controllers

import (
	"errors"
	"log"

	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/models"
)

func AddConversation(conversation models.Conversation) error {
	if len(conversation.Users) != 2 {
		return errors.New("conversation must contain 2 users and only 2 users")
	}
	result := initializers.DB.Create(conversation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindConversationByID(id string) (models.Conversation, error) {
	var conversation models.Conversation
	result := initializers.DB.Find(&conversation, "id = ?", id)
	if result.Error != nil {
		return conversation, result.Error
	}
	return conversation, nil
}

func FindConversationsByParticipants(userID1 string, userID2 string) {
	type ConversationsUsers struct {
		ConversationID string `json:"conversation_id"`
		UserID         string `json:"user_id"`
	}
	var conversationsUsers []ConversationsUsers
	err := initializers.DB.Raw("SELECT * from conversations_users WHERE user_id = ? OR user_id = ?", userID1, userID2).Scan(&conversationsUsers).Error
	if err != nil {
		log.Fatalf("err : %s", err.Error())
	}
	log.Println(len(conversationsUsers))
}

func FindAllConversations() ([]models.Conversation, error) {
	var conversations []models.Conversation
	result := initializers.DB.Find(&conversations)
	if result.Error != nil {
		return conversations, result.Error
	}
	return conversations, nil
}
