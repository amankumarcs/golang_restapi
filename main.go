package main

import (
	db "restapi/db"
	routes "restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Routes(router)
	db.ConnectDB()
	router.Run()
}

//err := json.Unmarshal([]byte(request.Body), &goldItems)
