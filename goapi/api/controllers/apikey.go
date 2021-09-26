package controllers

import (
	"goapi/api/service"
	"goapi/caching"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * IF NOT username IN memory
 * QUERY username, mac_addr FROM database
 * WHERE api_key LIKE postform
 * THEN caching IN memory
 */
type apiHandler struct {
	// Use service
	apiSrv service.APIkeyService
}

func NewAPIkeyHandler(apiSrv service.APIkeyService) apiHandler {
	return apiHandler{apiSrv: apiSrv}
}

// GET api key and caching in memory
func (h apiHandler) GetUserFromKey(ctx *gin.Context) {
	var p PostReceiver
	p.APIkey = ctx.PostForm("api_key") // api_key: ZG9sbHk7pJIhnO3ppHQvMS2-P9VR5XS2
	// fmt.Println("api key = ", p.APIkey)

	userFromKey, err := h.apiSrv.GetUserFromKey(p.APIkey)
	if err != nil {
		// Handle error ; api key not found in database
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusOK, userFromKey)

	// Caching in memory
	// Connect to redis
	redisClient := caching.ConnectRedis()

	// Call Setter - expire time = 60 second
	redisErr := caching.SetRedis(redisClient, userFromKey.Username, userFromKey.MacAdress, 60)
	if redisErr != nil {
		// Handle error
		log.Println("[can't set this key] ", redisErr)

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": redisErr.Error()})
		return
	}

	// After caching successfully
	log.Printf("SET %s %s\n", userFromKey.Username, userFromKey.MacAdress)

	// Call Getter - Check again username from POST == username from database
	_, redisErr = caching.GetRedis(redisClient, userFromKey.Username)
	if redisErr != nil {
		// Handle error
		log.Println("[key not found] ", redisErr)

		ctx.JSON(http.StatusNotFound, gin.H{
			"error":  redisErr.Error(),
			"detail": "username does not match",
		})

		return
	}

	// Response message if everything is OK
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Complete",
	})
}
