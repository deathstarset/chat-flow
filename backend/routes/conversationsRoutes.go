package routes

import (
	"github.com/deathstarset/backend-chatflow/handlers"
	"github.com/deathstarset/backend-chatflow/middlewares"
	"github.com/gin-gonic/gin"
)

func ConversationRoutes(g *gin.RouterGroup) {
	g.Use(middlewares.AuthOnly)
	g.POST("", handlers.CreateConversation)
	g.GET("", handlers.GetAllConversations)
	g.GET(":id", handlers.GetConversation)
	g.DELETE(":id", handlers.DeleteConversation)
}
