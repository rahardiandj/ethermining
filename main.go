package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rahardiandj/ethermining/src/api"
	"github.com/rahardiandj/ethermining/src/common"
	"github.com/rahardiandj/ethermining/src/model"
)

var (
	mongodbMod = model.MongoDBMod{}
	config     = &common.Configuration{}
)

func main() {

	config.Init()

	log.Printf("%+v", config)

	mongodbMod.Init(config)

	log.Printf("Initiating Service ... %v", config.App.Port)

	router := api.InitHandlers()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.App.Port), router))

}
