package db

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


type (
	FeatureProperties struct {
		Mapblklot string `bson:"MAPBLKLOT"`
		Blklot    string `bson:"BLKLOT"`
		BlockNum  string `bson:"BLOCK_NUM"`
	}

	Feature struct {
		Id         bson.ObjectId `bson:"_id,omitempty"`
		Kind       string `bson:"kind"`
		Properties FeatureProperties  `bson:"properties"`
	}
)

type (
	// UserController represents the controller for operating on the User resource
	FeatureController struct {
		session *mgo.Session
	}
)

var db *mgo.Database

func init() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("Feature")
}

func collection() *mgo.Collection {
	return db.C("FeatureCollection")
}

// GetAll returns all features from the database.
func GetAll() ([]Feature, error) {
	res := []Feature{}

	if err := collection().Find(nil).Limit(100).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}