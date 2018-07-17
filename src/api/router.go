package api

import (
	"github.com/gorilla/mux"
	"github.com/rahardiandj/ethermining/src/handler"
)

func InitHandlers() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/v1/api/bookmark/add/{tid}", handler.SetBookmarkHandler).Methods("POST")
	r.HandleFunc("/v1/api/bookmark/list", handler.GetBookmarkHandler).Methods("GET")
	r.HandleFunc("/v1/api/bookmark/remove/{tid}", handler.GetBookmarkHandler).Methods("DELETE")
	r.HandleFunc("/v1/api/bookmark/export", handler.ExportBookmarkHandler).Methods("POST")

	r.HandleFunc("/v1/api/blacklist/add/{tid}", handler.SetBlacklistHandler).Methods("POST")
	r.HandleFunc("/v1/api/blacklist/list", handler.GetBlacklistHandler).Methods("GET")
	r.HandleFunc("/v1/api/blacklist/remove/{tid}", handler.RemoveBlacklistHandler).Methods("DELETE")

	return r
}
