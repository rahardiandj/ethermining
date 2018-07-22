package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/rahardiandj/ethermining/src/common"
)

func (a *API) FetchTransaction(ctx context.Context) ([]common.Transaction, error) {

	var data = url.Values{}
	url := getFetchAPIURL()
	req, err := http.NewRequest("GET", url, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody := []common.Transaction{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &respBody)

	if err != nil {
		return nil, err
	}

	return respBody, err
}

func getFetchAPIURL() string {
	cfg := common.GetConfig()
	return fmt.Sprintf("%s?module=account&action=txlist&address=%s", cfg.APIEtherium.URL, cfg.APIEtherium.Address)

}
