package middlewares

import (
	"goapi/api/controllers"
	"goapi/caching"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p controllers.PostReceiver
		p.Username = c.PostForm("uname")
		// fmt.Println("username = ", p.Username)

		// Connect to redis
		redisClient := caching.ConnectRedis()

		// Caching ; Get
		value, redisErr := caching.GetRedis(redisClient, p.Username)
		if redisErr != nil {
			log.Println("[middlewares : key not found] ", redisErr)

			// Call Set
			c.Next() // Next to handler

		} else {
			log.Printf("test %v\n", value)
			c.JSON(http.StatusOK, gin.H{
				"message": "Ready",
			})
			c.Abort() // Stop
		}
	}
}
