package model

import (
	"log"
	"testing"

	"github.com/rahardiandj/ethermining/src/common"
)

var mongodbMod = MongoDBMod{}
var cfg = common.Configuration{}

func TestMain(m *testing.M) {
	cfg.InitTest()
	mongodbMod.Init(&cfg)
	log.Printf("Test 1")
}
