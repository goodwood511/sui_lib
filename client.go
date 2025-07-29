package sui_lib

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type TransactionResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Digest      string `json:"digest"`
		Transaction any    `json:"transaction"`
		Effects     any    `json:"effects"`
		Events      any    `json:"events"`
		TimestampMs string `json:"timestampMs"`
		Checkpoint  string `json:"checkpoint"`
	} `json:"result"`
}

func GetSuiTransactionBlock(txHash, rpcURL string) (*TransactionResponse, error) {
	client := resty.New()

	reqBody := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "sui_getTransactionBlock",
		"params": []any{
			txHash,
			map[string]bool{
				"showInput":          true,
				"showRawInput":       false,
				"showEffects":        true,
				"showEvents":         true,
				"showObjectChanges":  false,
				"showBalanceChanges": false,
				"showRawEffects":     false,
			},
		},
	}

	var response TransactionResponse
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetResult(&response).
		Post(rpcURL)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Sui RPC error: %s", resp.Status())
	}

	return &response, nil
}
