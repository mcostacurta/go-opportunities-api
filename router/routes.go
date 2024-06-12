package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/mcostacurta/gooportunities/docs"
	handler "github.com/mcostacurta/gooportunities/handler"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1Opennings := v1.Group("openings")

		v1Opennings.GET("/", addCommonResponseHeaders, handler.ListOpeningsHandler)
		v1Opennings.GET("/:id", addCommonResponseHeaders, handler.ShowOpeningHandler)
		v1Opennings.POST("/", addCommonResponseHeaders, handler.CreateOpeningHandler)
		v1Opennings.DELETE("/:id", addCommonResponseHeaders, handler.DeleteOpeningHandler)
		v1Opennings.PUT("/:id", addCommonResponseHeaders, handler.UpdateOpeningHandler)

	}
	router.GET("/health", addCommonResponseHeaders, handler.HealthCheck)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

func addCommonResponseHeaders(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache, no-store")
}
