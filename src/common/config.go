package common

import (
	"log"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	App      `json:"app"`
	Database `json:"database"`
	NSQ      `json:"nsq"`
	MongoDB  `json:"mongodb"`
}

type Database struct {
	Master string `json:"master"`
}

type App struct {
	Port int `json:"port"`
}

type NSQ struct {
	URL  string `json:"url"`
	Port int    `json:"port"`
}

type MongoDB struct {
	Master string `json:"master"`
	DBName string `json:"db_name"`
}

//Init Config
func (c *Configuration) Init() {

	config := Configuration{}
	err := gonfig.GetConf("config/ethermining.development.json", &config)

	if err != nil {
		log.Panic("Failed to initiate config  : ", err)
	}

	c.App = config.App
	c.Database = config.Database
	c.MongoDB = config.MongoDB
	c.NSQ = config.NSQ
}

func (c *Configuration) InitTest() {

	config := Configuration{}
	err := gonfig.GetConf("../../config/ethermining.test.json", &config)

	if err != nil {
		log.Panic("Failed to initiate config  : ", err)
	}

	c.App = config.App
	c.Database = config.Database
	c.MongoDB = config.MongoDB
	c.NSQ = config.NSQ
}
