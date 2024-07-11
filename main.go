package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PParist/go-otp-service/handler"
	"github.com/PParist/go-otp-service/repositories"
	"github.com/PParist/go-otp-service/service"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := initDatabase()
	redisRepository := repositories.NewRedisRepository(db)
	otpService := service.NewOtpEmailService(redisRepository)
	otpHandler := handler.NewMailOtpService(otpService)

	e := echo.New()
	e.POST("/sendotp", otpHandler.CreateMailOTP)
	e.GET("/validateotp/:uid", otpHandler.ValidateOTP)
	e.Logger.Fatal(e.Start(":8080"))
}

func initDatabase() *redis.Client {

	// เชื่อมต่อกับ Redis server
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), "6379"),
		//Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return redisClient
}
