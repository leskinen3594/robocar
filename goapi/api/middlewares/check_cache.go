package middlewares

import (
	"fmt"
	"goapi/api/forms"
	"goapi/caching"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var p forms.APIkeyReceiver

		// Bind request body - JSON
		if err := ctx.ShouldBindJSON(&p); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("username = ", p.Username)
		// fmt.Println("api key = ", p.APIkey)

		// Connect to redis
		redisClient := caching.NewConnectionRedis()

		// Caching ; Get
		value, redisErr := redisClient.GetRedis(p.Username)
		if redisErr != nil {
			// Error 'cause key (username) not found in memory
			log.Println("[middlewares : key not found] ", redisErr)

			// Set api_key variable
			ctx.Set("api_key", p.APIkey)

			// Next to handler
			ctx.Next()

		} else {
			log.Printf("GET %s %s\n", p.Username, value)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Ready",
				"robot":   value,
			})
			ctx.Abort() // Stop
		}
	}
}
