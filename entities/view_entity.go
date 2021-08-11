package entities

import "gopkg.in/mgo.v2/bson"

type View struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name   string        `bson:"name" json:"name"`
	UserID bson.ObjectId `bson:"userID" json:"userID"`
	Date   string        `bson:"date" json:"date"`
	Count  int           `bson:"count" json:"count"`
}
