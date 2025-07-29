package example

import (
	"github.com/goodwood511/sui_lib"
	"os"
	"testing"
)

func TestTransferSUI(t *testing.T) {
	privHex := os.Getenv("SUI_PRIVATE_KEY_HEX")
	if privHex == "" {
		t.Skip("SUI_PRIVATE_KEY_HEX env variable not set")
	}

	params := sui_lib.TxParams{
		RPCURL:    sui_lib.TestnetEndpointUrl,
		PkHex:     privHex,
		Recipient: os.Getenv("SUI_RECIPIENT_ADDRESS"),
		Amount:    1_000_000, // 0.001 SUI
	}

	digest, err := sui_lib.TransferSUI(params)
	if err != nil {
		t.Fatalf("TransferSUI failed: %v", err)
	}

	t.Logf("Transaction digest: %s", digest)
}
