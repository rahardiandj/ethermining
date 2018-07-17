package model

import (
	"log"

	"github.com/rahardiandj/ethermining/src/common"
	mgo "gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
)

type MongoDBMod struct {
	Session  interface{}
	Database *mgo.Database
}

func (m *MongoDBMod) Init(config *common.Configuration) {
	session, err := mgo.Dial(config.MongoDB.Master)

	log.Printf("Connecting to %v ...", config.MongoDB.Master)

	if err != nil {
		log.Panic("Failed to connect to mongodb server : ", err)
	}

	m.Session = session
	m.Database = session.DB(config.MongoDB.DBName)
	db = session.DB(config.MongoDB.DBName)
	log.Printf("Connected to %v - %v", config.MongoDB.Master, config.MongoDB.DBName)

}
