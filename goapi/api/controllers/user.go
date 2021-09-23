// Hexagonal Architecture ; Adapter connect to Service port
// REST Handlers ; UI Response
package controllers

import (
	"fmt"
	"goapi/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var w http.ResponseWriter

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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// Get a user filter by uid
func (h userHandler) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := h.usrSrv.GetUser(id)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
