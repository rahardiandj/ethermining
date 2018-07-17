package model

import (
	"context"
	"log"

	"github.com/rahardiandj/ethermining/src/common"
	"gopkg.in/mgo.v2/bson"
)

type BookmarkModel struct {
}

func InsertBookmark(ctx context.Context, bookmark common.Transaction) error {
	err := db.C(common.Bookmark).Insert(bookmark)
	if err != nil {
		log.Println(err)
	}
	return err

}

func GetBookmark(ctx context.Context) ([]common.Transaction, error) {
	var transactions = []common.Transaction{}

	err := db.C(common.Bookmark).Find(bson.M{}).All(&transactions)

	if err != nil {
		log.Println(err)
	}

	return transactions, err
}

func GetBookmarkByTxID(ctx context.Context, transID string) (*common.Transaction, error) {

	var bookmark = common.Transaction{}
	err := db.C(common.Bookmark).FindId(transID).One(&bookmark)
	return &bookmark, err
}

func RemoveBoomark(ctx context.Context, bookmark common.Transaction) error {

	err := db.C(common.Bookmark).Remove(&bookmark)

	if err != nil {
		log.Println(err)
	}
	return err
}
