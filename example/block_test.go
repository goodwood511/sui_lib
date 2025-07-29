package example

import (
	"github.com/goodwood511/sui_lib"
	"testing"
)

func TestGetLatestSuiBlockNumber(t *testing.T) {
	blockNumber, err := sui_lib.GetLatestSuiBlockNumber(sui_lib.TestnetEndpointUrl)
	if err != nil {
		t.Fatalf("Failed to get latest block number: %v", err)
	}
	t.Logf("Latest Sui Testnet Block Number: %s", blockNumber)
}
