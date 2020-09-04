package models

//Create Struct
type Tenants struct {
	ID     int 		`json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string   `json:"name,omitempty" bson:"name,omitempty"`
}