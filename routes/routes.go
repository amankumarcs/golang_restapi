package routes

import (
	controllers "restapi/controllers"
	middlewares "restapi/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/api", middlewares.DummyMiddleware, controllers.GetUsers)
	router.POST("/api", middlewares.DummyMiddleware, controllers.PostUsers)
	router.POST("/login", middlewares.DummyMiddleware, controllers.LoginUser)
	router.DELETE("/api", middlewares.DummyMiddleware, controllers.DeleteUsers)
}
