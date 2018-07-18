package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rahardiandj/ethermining/src/common"
	"github.com/rahardiandj/ethermining/src/model"
)

var transTest = common.Transaction{
	BlockNumber:        5176578,
	TimeStamp:          1519898336,
	Hash:               "0x05993680c24fc81909daab444d224c04ddad4ffbd514d5b24bc004ee801b1c6d",
	Nonce:              1,
	BlockHash:          "0xa5a1dc164508738eff07ad5bb00dcae1f3f2d304a70cffb7869bca70c37b67ee",
	TransactionIndex:   8,
	From:               "0x95343e65c188952ad41a56869b7cb6d89df8dd25",
	To:                 "0xfedae5642668f8636a11987ff386bfd215f942ee",
	Value:              0,
	Gas:                80137,
	GasPrice:           41000000000,
	IsError:            0,
	TxReceiptStatus:    1,
	Input:              "0xa9059cbb000000000000000000000000ab0f6b8c486fec656b270ff2b53ae09e454e12e10000000000000000000000000000000000000000000034511860c551606c0000",
	CommulativeGasUsed: 672102,
	GasUsed:            53425,
	Confirmations:      585807,
}

func SetBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	err := model.InsertBookmark(context.Background(), transTest)
	if err != nil {
		log.Printf("[Handler] [SetBookmarkHandler] %v\n", err)
	}
}

func GetBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	bookmarks, err := model.GetBookmark(context.Background())
	if err != nil {
		log.Printf("[Handler] [GetBookmarkHandler] %v\n", err)
	}

	bodyJson, err := json.Marshal(bookmarks)

	if err != nil {
		log.Printf("[Handler] [GetBookmarkHandler] %v\n", err)
	}

	w.Write(bodyJson)
}

func RemoveBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var transID string
	bookmark, err := model.GetBookmarkByTxID(ctx, transID)

	if err != nil {
		log.Printf("[Handler] [RemoveBookmarkHandler] %v\n", err)
	}

	err = model.RemoveBoomark(ctx, *bookmark)

	if err != nil {
		log.Printf("[Hanlder] [RemoveBookmarkHandler] %v\n", err)
	}
}

func ExportBookmarkHandler(w http.ResponseWriter, r *http.Request) {

}
