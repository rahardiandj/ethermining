package handler_test

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

type recFunc func(*httptest.ResponseRecorder) error

var recFuncs = func(fns ...recFunc) []recFunc { return fns }

var checkStatus = func(expect int) recFunc {
	return func(rec *httptest.ResponseRecorder) error {
		if rec.Code != expect {
			return fmt.Errorf("Expect http code %v, return %v", expect, rec.Code)
		}
		return nil
	}
}

var checkHeader = func(key, expect string) recFunc {
	return func(rec *httptest.ResponseRecorder) error {
		if result := rec.Result().Header.Get(key); result != expect {
			return fmt.Errorf("Expect header %v, return %v", expect, result)
		}
		return nil
	}
}

var checkBody = func(expect string) recFunc {
	return func(rec *httptest.ResponseRecorder) error {
		if contain := rec.Body.String(); contain != expect {
			return fmt.Errorf("Expect body %v, return %v", expect, contain)
		}
	}
}

func TestMain(m *testing.Main) {

}
