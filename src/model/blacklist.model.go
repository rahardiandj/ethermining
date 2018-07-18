package model

import (
	"context"
	"log"

	"github.com/rahardiandj/ethermining/src/common"
	"gopkg.in/mgo.v2/bson"
)

//InsertBlacklist is to Insert Blacklisted Transaction to Documents
func InsertBlacklist(ctx context.Context, blacklist common.Transaction) error {
	err := db.C(common.BlacklistDocName).Insert(&blacklist)

	if err != nil {
		log.Println("[Model] [InsertBlacklist] : ", err)
	}

	return err
}

//GetBlacklist is to Get All Blacklisted Transaction from Documents
//TODO Limit the number transaction and include filtering
func GetBlacklist(ctx context.Context) ([]common.Transaction, error) {
	var transactions = []common.Transaction{}

	err := db.C(common.BlacklistDocName).Find(bson.M{}).All(&transactions)

	if err != nil {
		log.Println("[Model] [GetBlacklist] : ", err)
	}

	return transactions, err
}

//GetBlacklistByTxID is to Get Blacklisted Transaction by hash code
func GetBlacklistByTxID(ctx context.Context, transID string) (*common.Transaction, error) {

	var bookmark = common.Transaction{}
	err := db.C(common.BookmarkDocName).FindId(transID).One(&bookmark)

	if err != nil {
		log.Println("[Model] [GetBookmarkByTxID] : ", err)
	}

	return &bookmark, err
}

//RemoveBlacklist is to Get Blacklisted Transaction
func RemoveBlacklist(ctx context.Context, blacklist common.Transaction) error {
	err := db.C(common.BlacklistDocName).Remove(&blacklist)

	if err != nil {
		log.Println("[Model] [RemoveBlacklist] : ", err)
	}
	return err
}
