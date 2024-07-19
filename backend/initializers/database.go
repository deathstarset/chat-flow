package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RD *redis.Client

func ConnectDb() {
	var err error
	dbUrl := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Println("Connected Succefully to database")
}

func InitRedis() {
	redisPass := os.Getenv("REDIS_PASSWORD")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	RD = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPass,
		DB:       0,
	})

	err := RD.Ping().Err()
	if err != nil {
		log.Fatalf("Failed to connect to redis : %s", err.Error())
	}
	log.Println("Connected Succefully to redis")
}
