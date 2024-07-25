package controllers

import (
	"errors"

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
	result := initializers.DB.Preload("Users").Find(&conversation, "id = ?", id)
	if result.Error != nil {
		return conversation, result.Error
	}
	return conversation, nil
}

func FindConversationsByParticipants(userID1 string, userID2 string) (int, error) {
	type ConversationsUsers struct {
		ConversationID string `json:"conversation_id"`
		UserID         string `json:"user_id"`
	}
	var conversationsUsers []ConversationsUsers
	err := initializers.DB.Raw("SELECT * from conversations_users WHERE user_id = ? OR user_id = ?", userID1, userID2).Scan(&conversationsUsers).Error
	if err != nil {
		return len(conversationsUsers), err
	}
	return len(conversationsUsers), nil
}

func FindAllConversations() ([]models.Conversation, error) {
	var conversations []models.Conversation
	result := initializers.DB.Find(&conversations)
	if result.Error != nil {
		return conversations, result.Error
	}
	return conversations, nil
}

func FindConversationsByUserID(userID string) ([]models.Conversation, error) {
	var user models.User
	result := initializers.DB.Preload("Conversations").Find(&user, "id = ?", userID)
	if result.Error != nil {
		return user.Conversations, result.Error
	}
	return user.Conversations, nil
}

func RemoveConversation(id string) error {
	var conversation models.Conversation

	result := initializers.DB.Where("id = ?", id).First(&conversation)
	if result.Error != nil {
		return result.Error
	}

	initializers.DB.Model(&conversation).Association("Users").Clear()

	result = initializers.DB.Delete(&models.Conversation{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
