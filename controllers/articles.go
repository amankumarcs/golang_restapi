package articles

import (
	"fmt"
	"context"
	"github.com/gin-gonic/gin"
	db "restapi/db"
	models "restapi/models"
	"go.mongodb.org/mongo-driver/bson"
)


func GetUsers(c *gin.Context) {
	collection := db.GetCollection("test");
	fmt.Println("sending data",collection)
	var tenants models.Tenants;

	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&tenants)
	fmt.Println("err data",err)
	if err != nil {
		return
	}
	fmt.Println("err data",tenants.Name)

   c.JSON(200, gin.H{"data" : tenants})
} 