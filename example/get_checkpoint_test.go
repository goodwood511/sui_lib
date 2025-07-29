package example

import (
	"github.com/goodwood511/sui_lib"
	"testing"
)

func TestGetCheckpointTransactions(t *testing.T) {
	checkpointNumber := "223780535"

	txs, err := sui_lib.GetCheckpointTransactions(sui_lib.TestnetEndpointUrl, checkpointNumber)
	if err != nil {
		t.Fatalf("Failed to get checkpoint transactions: %v", err)
	}

	if len(txs) == 0 {
		t.Errorf("Expected transactions, got none")
	}

	t.Logf("Retrieved %d transactions. First: %s", len(txs), txs[0])
}
