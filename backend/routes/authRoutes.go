package routes

import (
	"github.com/deathstarset/backend-chatflow/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(g *gin.RouterGroup) {
	g.POST("login", handlers.Login)
	g.POST("register", handlers.Register)
}
