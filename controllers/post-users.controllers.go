package controllers

import (
	"context"
	"log"
	"net/http"
	db "restapi/db"
	models "restapi/models"

	"github.com/gin-gonic/gin"
)

func PostUsers(c *gin.Context) {
	collection := db.GetCollection("customers")

	var userData models.Tenants

	err := c.ShouldBindJSON(&userData)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	result, err := collection.InsertOne(context.TODO(), userData)
	if err != nil {
		log.Printf("Could not create user: %v", err)
	}
	c.JSON(200, gin.H{"statusCode": 200, "message": "customer found", "data": result})
}
