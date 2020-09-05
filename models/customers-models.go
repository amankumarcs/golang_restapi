package models

//Create Struct
type Tenants struct {
	ID       int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
