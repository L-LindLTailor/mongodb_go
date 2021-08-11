package entities

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name" json:"name"`
}
