package routes

import (
	"github.com/deathstarset/backend-chatflow/handlers"
	"github.com/deathstarset/backend-chatflow/middlewares"
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(g *gin.RouterGroup) {

	g.Use(middlewares.AuthOnly)
	g.GET("", handlers.GetProfile)
	g.PUT("", handlers.UpdateProfile)
	g.DELETE("", handlers.DeleteProfile)
}
