package common

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {

	var dbconf = Database{
		Master: "localhost:3360",
	}

	var appconf = App{
		Port: 8080,
	}

	var nsqconf = NSQ{
		URL:  "http://localhost",
		Port: 4432,
	}

	var mongodbconf = MongoDB{
		Master: "http://localhost:27017",
	}
	var config = Configuration{
		App:      appconf,
		Database: dbconf,
		NSQ:      nsqconf,
		MongoDB:  mongodbconf,
	}

	configJson, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(configJson))

}
