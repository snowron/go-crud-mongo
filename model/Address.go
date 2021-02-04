package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Latitute  float64 `bson:"Latitude" json:"Latitude"`
	Longitude float64 `bson:"Longitude" json:"Longitude"`
}

type Address struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Country  string             `bson:"Country" json:"Country"`
	City     string             `bson:"City" json:"City"`
	FullText string             `bson:"FullText" json:"FullText"`
	Location Location           `bson:"Location" json:"Location"`
}
