package routes

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/controllers"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	router.GET("/auth", controllers.AuthAction)

	router.Run(configs.GetEnv().AuthServiceBaseUrl)
	return router
}
