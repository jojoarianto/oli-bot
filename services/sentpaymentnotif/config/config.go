package config

import (
	"gopkg.in/mgo.v2"
	"log"
)

// Connect make connection
// please dont forget to db.Close()
func Connect(strUri, dbName string) *mgo.Database {
	session, err := mgo.Dial(strUri)
	if err != nil {
		log.Fatal(err)
	}
	return session.DB(dbName)
}
