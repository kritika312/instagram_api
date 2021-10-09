package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	ID     primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"Name,omitempty" bson:"isbn,omitempty"`
	Email  string             `json:"Email" bson:"title,omitempty"`
	Password  string            `json:"Paasword" bson:"author,omitempty"`
}

type Posts struct {
    ID     primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`
	caption  string             `json:"caption,omitempty" bson:"isbn,omitempty"`
	ImageURL  string             `json:"ImageUrl" bson:"title,omitempty"`
	PostedTimestamp  string            `json:"PostedTimestamp" bson:"author,omitempty"`
	
}
