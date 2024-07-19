package main

import (
	"os"

	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/routes"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDb()
	initializers.InitRedis()

}
func main() {
	r := gin.Default()
	port := os.Getenv("PORT")

	// init redis store
	// configuring cors
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	cfg.AllowHeaders = []string{"Accept", "Authorization"}
	r.Use(cors.New(cfg))

	// serving static images
	r.Static("/api/v1/uploads/profile", "./uploads/profile")
	routes.UserRoutes(r.Group("/api/v1/users"))
	routes.AuthRoutes(r.Group("/api/v1/auth"))
	routes.ProfileRoutes(r.Group("/api/v1/profile"))

	r.Run(port)
}
