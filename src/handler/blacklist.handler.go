package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rahardiandj/ethermining/src/model"
)

func SetBlacklistHandler(w http.ResponseWriter, r *http.Request) {

	err := model.InsertBookmark(context.Background(), transTest)
	if err != nil {
		log.Printf("%v", err)
	}
}

func GetBlacklistHandler(w http.ResponseWriter, r *http.Request) {
	bookmarks, err := model.GetBookmark(context.Background())
	if err != nil {
		log.Printf("%v", err)
	}

	bodyJson, err := json.Marshal(bookmarks)

	if err != nil {
		log.Printf("%v", err)
	}

	w.Write(bodyJson)
}

func RemoveBlacklistHandler(w http.ResponseWriter, r *http.Request) {

}
