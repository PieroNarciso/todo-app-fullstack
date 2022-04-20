package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
    Id primitive.ObjectID `bson:"_id" json:"id,omitempty"`
    Title string `bson:"title" json:"title,omitempty"`
    Description string `bson:"description" json:"description,omitempty"`
    Completed bool `bson:"completed" json:"completed"`
}
