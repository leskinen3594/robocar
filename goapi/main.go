package main

import (
	"fmt"
	"goapi/api/controllers"
	"goapi/api/middlewares"
	"goapi/api/models"
	"goapi/api/service"
	"goapi/caching"
	"goapi/configs"
	"goapi/messagebroker"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Connect to database
	db, err := configs.ConnectDB()
	if err != nil {
		log.Println("Database error: ", err)
	}

	// Connect to MQTT
	mqttConnection := messagebroker.NewConnectionMQTT()
	go mqttConnection.Subscribe("/mR_robot/ws/#")

	// Connect to Redis
	redisConnction := caching.NewConnectionRedis()

	// Users Entity
	userRepo := models.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	userHandler := controllers.NewUserHandler(userService)

	// APIkeys Entity
	apiRepo := models.NewAPIkeyRepositoryDB(db)
	apiService := service.NewAPIkeyService(apiRepo)
	apiHandler := controllers.NewAPIkeyHandler(apiService)

	// Robot Handler
	robotHandler := controllers.NewRobotHandler(mqttConnection, redisConnction)

	// Router - endpoints
	// GET		method	- GetUserAll, GetUserByID, ...
	// POST		method	- NewUser, CreateAPIkey, CheckAPIkey, ...
	// PUT		method	- ...
	// DELETE	method	- ...
	// Group for request interacted w/ database
	// router.Use(middlewares.Logger())	// Use custom logger
	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/users", userHandler.GetUsers)
		apiRoute.GET("/users/:id", userHandler.GetUser)
		apiRoute.Use(middlewares.CheckCache()).POST("/handshake", apiHandler.GetUserFromKey)
	}

	// Group for request interacted w/ robot
	robotRoute := router.Group("/robot")
	{
		robotRoute.POST("/handshake", robotHandler.Handshake) // Publish to robot
		robotRoute.POST("/movement", robotHandler.Movement)
	}

	// Run server
	host, port := configs.AppConfig()
	router.Run(fmt.Sprintf("%v:%v", host, port)) // Server running default on 0.0.0.0:8080
}
