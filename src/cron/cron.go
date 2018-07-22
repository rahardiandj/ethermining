package main

import (
	"log"

	"github.com/rahardiandj/ethermining/src/common"
	"github.com/robfig/cron"
)

func main() {
	var config = &common.Configuration{}

	config.Init()

	c := cron.New()
	c.AddFunc(config.CronFetchTransaction.Interval, FetchTransactionList)

	c.Start()
	log.Println("Cron Started ... ")
}

//TODO Add Publish to MessageQueue after fetching
func FetchTransactionList() {

}
