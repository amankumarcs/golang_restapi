package controllers

import (
	"context"
	"fmt"
	"log"
	db "restapi/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *gin.Context) {
	collection := db.GetCollection("customers")

	fmt.Println("sending data 1")

	filterCustomers, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var users []bson.M

	if err = filterCustomers.All(context.TODO(), &users); err != nil {
		log.Fatal(err.Error())
	}

	if err != nil {
		return
	}

	c.JSON(200, gin.H{"statusCode": 200, "message": "customer found", "data": users})
}
