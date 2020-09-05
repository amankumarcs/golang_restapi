package helpers

import "go.mongodb.org/mongo-driver/bson"

var ErrorMessages = bson.M{
	"StatusUnauthorized":        "please provide valid login details.",
	"LoginSuccess":              "Login successfully.",
	"StatusUnprocessableEntity": "oops something went wrong.",
}

type Error struct {
	statusCode int
	message    string
}
