package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	db "restapi/db"
	"restapi/helpers"
	middleware "restapi/middlewares"
	models "restapi/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginUser(c *gin.Context) {
	collection := db.GetCollection("customers")

	var requestData Login

	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"statusCode": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}
	query := bson.M{"email": requestData.Email}

	var userData models.Tenants

	err = collection.FindOne(context.TODO(), query).Decode(&userData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"statusCode": http.StatusUnauthorized, "message": helpers.ErrorMessages["StatusUnauthorized"]})
		return
	}

	if requestData.Password != userData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"statusCode": http.StatusUnauthorized, "message": helpers.ErrorMessages["StatusUnauthorized"]})
		return
	}

	token, err := middleware.GenerateToken(userData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(200, gin.H{"statusCode": 200, "message": helpers.ErrorMessages["LoginSuccess"], "data": bson.M{"token": token, "_id": userData.ID, "name": userData.Name}})
}
