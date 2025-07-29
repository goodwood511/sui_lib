package example

import (
	"github.com/goodwood511/sui_lib"
	"testing"
)

func TestGetSuiTransactionBlock(t *testing.T) {
	txHash := "GEkgLs1LVb2MHYPtoZrK7ZzGD9NjxP7hk2gvHH56G5q"

	result, err := sui_lib.GetSuiTransactionBlock(txHash, sui_lib.TestnetEndpointUrl)
	if err != nil {
		t.Fatalf("GetSuiTransactionBlock error: %v", err)
	}

	if result.Result.Digest != txHash {
		t.Errorf("Expected digest %s, got %s", txHash, result.Result.Digest)
	}

	t.Logf("Digest: %s", result.Result.Digest)
	t.Logf("TimestampMs: %s", result.Result.TimestampMs)
	t.Logf("Checkpoint: %s", result.Result.Checkpoint)
}
