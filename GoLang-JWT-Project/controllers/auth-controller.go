package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response("msg : ", "Auth Service Run"))
}

func response(key string, value string) gin.H {
	return gin.H{
		key: value,
	}
}
