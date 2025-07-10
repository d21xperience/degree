package routes

import (
	"compile-service/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/cs/compile", controllers.CompileHandler)
	r.GET("/cs/version", controllers.VersionCompilerHandler)
}
