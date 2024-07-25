package main

import (
	"log"

	"github.com/deathstarset/backend-chatflow/initializers"
	"github.com/deathstarset/backend-chatflow/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDb()
}
func main() {
	var exists bool
	initializers.DB.Raw("SELECT EXISTS(SELECT FROM pg_type WHERE typname = 'role_enum');").Scan(&exists)
	if !exists {
		result := initializers.DB.Exec("CREATE TYPE role_enum AS ENUM ('Normal', 'Admin');")
		if result.Error != nil {
			log.Fatalf("Failed to create enum type: %v", result.Error.Error())
		}
	}
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Conversation{})
	log.Println("Migration success")
}
