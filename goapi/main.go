package main

import (
	"fmt"
	"goapi/api/controllers"
	"goapi/api/middlewares"
	"goapi/api/models"
	"goapi/api/service"
	"goapi/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database
	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Users Entity
	userRepo := models.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	userHandler := controllers.NewUserHandler(userService)

	// APIkeys Entity
	apiRepo := models.NewAPIkeyRepositoryDB(db)
	apiService := service.NewAPIkeyService(apiRepo)
	apiHandler := controllers.NewAPIkeyHandler(apiService)

	/**
	 * Router - endpoints
	 * GET		method	- GetUserAll, GetUserByID, ...
	 * POST		method	- NewUser, CreateAPIkey, CheckAPIkey, ...
	 * PUT		method	- ...
	 * DELETE	method	- ...
	 */
	// Group for request interacted w/ database
	router.Use(middlewares.Logger())
	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/users", userHandler.GetUsers)
		apiRoute.GET("/users/:id", userHandler.GetUser)
		apiRoute.Use(middlewares.CheckCache()).POST("/handshake", apiHandler.GetUserFromKey)
	}

	// Group for request interacted w/ robot
	// robotRoute := router.Group("/robot")
	// {
	// 	robotRoute.POST("/handshake", )
	// }

	// Run server
	host, port := configs.AppConfig()
	router.Run(fmt.Sprintf("%v:%v", host, port)) // server running default on 0.0.0.0:8080
}
