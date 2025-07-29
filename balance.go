package sui_lib

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SuiBalanceParams struct {
	Address  string
	CoinType string
}

type SuiBalanceResult struct {
	CoinType        string `json:"coinType"`
	CoinObjectCount int    `json:"coinObjectCount"`
	TotalBalance    string `json:"totalBalance"`
	LockedBalance   any    `json:"lockedBalance"`
}

type SuiResponse struct {
	Jsonrpc string           `json:"jsonrpc"`
	ID      int              `json:"id"`
	Result  SuiBalanceResult `json:"result"`
}

func GetSuiBalance(params SuiBalanceParams, rpcURL string) (*SuiBalanceResult, error) {
	client := resty.New()

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "suix_getBalance",
		"params":  []interface{}{params.Address, params.CoinType},
	}

	var response SuiResponse
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&response).
		Post(rpcURL)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status())
	}

	return &response.Result, nil
}
