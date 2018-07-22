package model

import (
	"context"
	"log"

	"github.com/rahardiandj/ethermining/src/common"
	"github.com/rahardiandj/ethermining/src/constant"
	"gopkg.in/mgo.v2/bson"
)

//InsertBookmark is to Insert Bookmarked Transaction to Documents
func InsertBookmark(ctx context.Context, bookmark common.Transaction) error {
	err := db.C(constant.BookmarkDocName).Insert(bookmark)
	if err != nil {
		log.Println("[Model] [InsertBookmark] : ", err)
	}
	return err

}

//GetBookmark is to Get All Bookmark Transaction from Documents
//TODO Limit the number transaction and include filtering
func GetBookmark(ctx context.Context) ([]common.Transaction, error) {
	var transactions = []common.Transaction{}

	err := db.C(constant.BookmarkDocName).Find(bson.M{}).All(&transactions)

	if err != nil {
		log.Println("[Model] [GetBookmark] : ", err)
	}

	return transactions, err
}

//GetBookmarkByTxID is to Get Bookmarked Transaction by hash code
func GetBookmarkByTxID(ctx context.Context, transID string) (*common.Transaction, error) {

	var bookmark = common.Transaction{}
	err := db.C(constant.BookmarkDocName).FindId(transID).One(&bookmark)

	if err != nil {
		log.Println("[Model] [GetBookmarkByTxID] : ", err)
	}

	return &bookmark, err
}

//RemoveBoomark is to Remove Blacklisted Transaction
func RemoveBoomark(ctx context.Context, bookmark common.Transaction) error {

	err := db.C(constant.BookmarkDocName).Remove(&bookmark)

	if err != nil {
		log.Println("[Model] [RemoveBoomark] : ", err)
	}
	return err
}
