package routes

import (
	"github.com/deathstarset/backend-chatflow/handlers"
	"github.com/deathstarset/backend-chatflow/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.RouterGroup) {

	g.Use(middlewares.AdminOnly)
	g.POST("", handlers.CreateUser)
	g.GET("", handlers.GetAllUsers)
	g.GET(":id", handlers.GetUser)
	g.PUT(":id", handlers.UpdateUser)
	g.DELETE(":id", handlers.DeleteUser)
}
