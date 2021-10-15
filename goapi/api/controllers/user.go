// Hexagonal Architecture ; Adapter connect to Service port
// REST Handlers ; UI Response
package controllers

import (
	"goapi/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	usrSrv service.UserService
}

func NewUserHandler(usrSrv service.UserService) userHandler {
	return userHandler{usrSrv: usrSrv}
}

// Get all user
func (h userHandler) GetUsers(ctx *gin.Context) {
	users, err := h.usrSrv.GetUsers()
	if err != nil {
		// Handle error
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// Get a user filter by uid
func (h userHandler) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := h.usrSrv.GetUser(id) // Get Body
	if err != nil {
		// Handle error
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
