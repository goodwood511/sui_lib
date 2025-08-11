package sui_lib

import (
	"context"
	"encoding/json"
	"fmt"

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

func GetCheckpointTransactions(rpcURL string, checkpointNumber string) (*CheckpointResponse, string, error) {
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
		return nil, "", err
	}
	if resp.StatusCode() != 200 {
		return nil, string(resp.Body()), fmt.Errorf("unexpected http status code: %d", resp.StatusCode())
	}
	bodyBytes := resp.Body()
	bodyString := string(bodyBytes)

	var result CheckpointResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		// 解析失败时，返回原始body，方便调试
		return nil, bodyString, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return &result, bodyString, nil
}
