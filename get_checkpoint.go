package sui_lib

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type CheckpointResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Epoch                    string   `json:"epoch"`
		SequenceNumber           string   `json:"sequenceNumber"`
		Digest                   string   `json:"digest"`
		NetworkTotalTransactions string   `json:"networkTotalTransactions"`
		PreviousDigest           string   `json:"previousDigest"`
		TimestampMs              string   `json:"timestampMs"`
		Transactions             []string `json:"transactions"`
	} `json:"result"`
}

func GetCheckpointTransactions(rpcURL string, checkpointNumber string) ([]string, error) {
	client := resty.New()

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "sui_getCheckpoint",
		"params":  []interface{}{checkpointNumber},
	}

	resp, err := client.R().
		SetContext(context.Background()).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(rpcURL)

	if err != nil {
		return nil, err
	}

	var result CheckpointResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	return result.Result.Transactions, nil
}
