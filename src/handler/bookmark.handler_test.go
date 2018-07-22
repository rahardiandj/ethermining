package handler_test

import (
	"net/http"
	"testing"
)

//Testing Get BlackList Handler using recorder
func TestGetBookmarkHandler(t *testing.T) {

	tests := map[string]struct {
		h    func(w http.ResponseWriter, r *http.Request)
		recs []recFunc
	}{
		"testBookmarktHandlerSuccess": {
			func(w http.ResponseWriter, r *http.Request) {},
			recFuncs(checkStatus(200)),
		},
		{
			"testBlookmarktContent": {
				func(w http.ResponseWriter, r *http.Request) {},
				recFunc(checkStatus(201), checkHeader("Content-Type", "application/json"), checkBody("hash")),
			},
		},
	}

	r, _ := http.NewRequest("GET", "http://localhost:9090/v1/api/bookmark/list", nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := http.HandlerFunc(tt.h)
			rec := NewRecorder()
			h.ServeHTTP(rec, r)
			for _, check := range tt.checks {
				if err := check(rec); err != nil {
					t.Error(err)
				}
			}
		})
	}

}
