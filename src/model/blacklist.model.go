package model

import (
	"context"
	"log"

	"github.com/rahardiandj/ethermining/src/common"
	"gopkg.in/mgo.v2/bson"
)

type BlacklistModel struct {
}

func InsertBlacklist(ctx context.Context, blacklist common.Transaction) error {
	err := db.C(common.Blacklist).Insert(&blacklist)
	return err
}

func GetBlacklist(ctx context.Context) ([]common.Transaction, error) {
	var transactions = []common.Transaction{}

	err := db.C(common.Bookmark).Find(bson.M{}).All(&transactions)

	if err != nil {
		log.Println(err)
	}

	return transactions, err
}

func RemoveBlacklist(ctx context.Context, blacklistID string) error {

	return nil
}
