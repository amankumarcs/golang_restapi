package main

import (
	"github.com/gin-gonic/gin"
	db "restapi/db"
    routes "restapi/routes"
)


func main(){
	router := gin.Default()
	routes.Routes(router)
	db.ConnectDB();
	router.Run()
}


//err := json.Unmarshal([]byte(request.Body), &goldItems)


