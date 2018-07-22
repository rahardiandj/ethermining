package common

import (
	"log"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	App         `json:"app"`
	Database    `json:"database"`
	NSQ         `json:"nsq"`
	MongoDB     `json:"mongodb"`
	APIEtherium `json:"api_etherium"`
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

type APIEtherium struct {
	URL     string `json:"url"`
	Address string `json:"address"`
}

var config = Configuration{}

//Init Config
func (c *Configuration) Init() {

	err := gonfig.GetConf("config/ethermining.development.json", &config)

	if err != nil {
		log.Panic("Failed to initiate config  : ", err)
	}

	c.App = config.App
	c.Database = config.Database
	c.MongoDB = config.MongoDB
	c.NSQ = config.NSQ

}

func GetConfig() *Configuration {
	return &config
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
	c.APIEtherium = config.APIEtherium
}
