package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	ID     primitive.ObjectID `json:"ID,omitempty" bson:"ID,omitempty"`
	Name   string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Email  string             `json:"Email" bson:"Email,omitempty"`
	Password  string            `json:"Paasword" bson:"Password,omitempty"`
}

type Posts struct {
    ID     primitive.ObjectID `json:"ID,omitempty" bson:"ID,omitempty"`
	caption  string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL  string             `json:"ImageUrl" bson:"ImageURL,omitempty"`
	PostedTimestamp  string            `json:"PostedTimestamp" bson:"PostedTimestamp,omitempty"`
	
}
