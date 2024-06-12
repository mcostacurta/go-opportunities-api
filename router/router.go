package router

import "github.com/gin-gonic/gin"

func Initialize() *gin.Engine {
	router := gin.Default()

	initializeRoutes(router)

	return router
}
