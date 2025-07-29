package sui_lib

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type SuiRPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type SuiRPCResponseLatestSuiBlockNumber struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

func GetLatestSuiBlockNumber(rpcURL string) (string, error) {
	client := resty.New()

	reqBody := SuiRPCRequest{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "sui_getLatestCheckpointSequenceNumber",
		Params:  []interface{}{},
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Post(rpcURL)
	if err != nil {
		return "", err
	}

	var rpcResp SuiRPCResponseLatestSuiBlockNumber
	if err := json.Unmarshal(resp.Body(), &rpcResp); err != nil {
		return "", err
	}

	return rpcResp.Result, nil
}
