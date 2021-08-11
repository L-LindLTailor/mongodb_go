package config

import "gopkg.in/mgo.v2"

func GetMongoDB() (*mgo.Database, error) {
	host := "mongodb://localhost:27017"
	dbname := "First_Database"

	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	db := session.DB(dbname)
	return db, err
}
