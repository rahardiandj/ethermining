package model

import (
	"context"
	"testing"

	"github.com/rahardiandj/ethermining/src/common"
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

//Test Bookmark Directly To Database
func TestCRUDBookmark(t *testing.T) {
	err := InsertBookmark(context.Background(), transTest)
	if err != nil {
		t.Log("[TestCRUDBookmark] [InsertBookmark]", err)
	}

	bookmark, err := GetBookmarkByTxID(context.Background(), transTest.Hash)

	if err != nil {
		t.Log("[TestCRUDBookmark] [GetBookmarkByTxID]", err)
	}

	err = RemoveBoomark(context.Background(), *bookmark)

	if err != nil {
		t.Log("[TestCRUDBookmark] [RemoveBoomark]", err)
	}

}

//Test Blacklist Directly to Database
func TestCRUDBlacklist(t *testing.T) {
	err := InsertBlacklist(context.Background(), transTest)
	if err != nil {
		t.Log("[TestCRUDBookmark] [Insert]", err)
	}

	blacklist, err := GetBlacklistByTxID(context.Background(), transTest.Hash)

	if err != nil {
		t.Log("[TestCRUDBookmark] [Insert]", err)
	}

	err = RemoveBlacklist(context.Background(), *blacklist)

	if err != nil {
		t.Log("[TestCRUDBookmark] [RemoveBlacklist]", err)
	}
}
