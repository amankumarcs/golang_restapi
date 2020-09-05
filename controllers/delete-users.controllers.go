package controllers

import (
	"context"
	"log"
	db "restapi/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type users struct {
	Ids []int `json:"ids,omitempty" bson:"ids,omitempty"`
}

func DeleteUsers(c *gin.Context) {
	collection := db.GetCollection("customers")

	var userIds users
	if err := c.BindJSON(&userIds); err != nil {
		return
	}

	query := bson.D{{"_id", bson.D{{"$in", userIds.Ids}}}}
	result, err := collection.DeleteMany(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount == 0 {
		c.JSON(200, gin.H{"statusCode": 200, "message": "customer not deleated"})
		return
	}

	c.JSON(200, gin.H{"statusCode": 200, "message": "customer successfully deleted", "data": result})
}
