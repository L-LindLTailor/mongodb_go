package entities

import "gopkg.in/mgo.v2/bson"

type Product struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Price    float32       `bson:"price" json:"price"`
	Quantity int           `bson:"quantity" json:"quantity"`
	Status   bool          `bson:"status" json:"status"`
}
