package routes

import 	(
	"github.com/gin-gonic/gin"
	middlewares "restapi/middlewares"
	articles "restapi/controllers"
)

func Routes(router *gin.Engine){
	router.GET("/api", middlewares.DummyMiddleware, articles.GetUsers)
}