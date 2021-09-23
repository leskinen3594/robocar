package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) error {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	return nil
}
