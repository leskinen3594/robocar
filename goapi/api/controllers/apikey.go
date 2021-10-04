package controllers

import (
	"fmt"
	"goapi/api/service"
	"goapi/caching"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IF NOT username IN memory
// QUERY username, mac_addr FROM database
// WHERE api_key LIKE postform
// THEN caching IN memory
type apiHandler struct {
	// Use service
	apiSrv service.APIkeyService
}

func NewAPIkeyHandler(apiSrv service.APIkeyService) apiHandler {
	return apiHandler{apiSrv: apiSrv}
}

// Query username, mac address form api key and caching in memory
func (h apiHandler) GetUserFromKey(ctx *gin.Context) {
	// api_key: ZG9sbHk7pJIhnO3ppHQvMS2-P9VR5XS2

	// Get variable from middlewares
	key := ctx.MustGet("api_key").(string)
	fmt.Println("api key = ", key)
	uname := ctx.MustGet("uname").(string)
	fmt.Println("uname = ", uname)

	userFromKey, err := h.apiSrv.GetUserFromKey(key, uname)
	if err != nil {
		// Handle error ; api key not found in database
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// Caching in memory
	// Connect to redis
	redisClient := caching.NewConnectionRedis()

	// Call Setter - expire time = 3,600 second (1 hour)
	redisSetErr := redisClient.SetRedis(userFromKey.Username, userFromKey.MacAdress, 3600)
	if redisSetErr != nil {
		// Handle error
		log.Println("[can't set this key] ", redisSetErr)
	} else {
		// Caching successfully
		log.Printf("SET %s %s\n", userFromKey.Username, userFromKey.MacAdress)
	}

	// Call Getter - Check again username from POST == username from database
	val, redisGetErr := redisClient.GetRedis(userFromKey.Username)
	if redisGetErr != nil {
		// Handle error
		log.Println("[key not found] ", redisGetErr)

		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "username does not match",
			"success": false,
		})

		return
	} else {
		// Response message if everything is OK
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Connected",
			"robot":   val,
			"success": true,
		})
	}
}
