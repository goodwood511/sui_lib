package example

import (
	"github.com/goodwood511/sui_lib"
	"os"
	"testing"
)

func TestComposeTx(t *testing.T) {
	privHex := os.Getenv("SUI_PRIVATE_KEY_HEX")
	if privHex == "" {
		t.Skip("SUI_PRIVATE_KEY_HEX env variable not set")
	}

	params := sui_lib.TxParams{
		RPCURL:    sui_lib.TestnetEndpointUrl,
		PkHex:     privHex,
		Owner:     os.Getenv("SUI_OWNER_ADDRESS"),
		Recipient: os.Getenv("SUI_RECIPIENT_ADDRESS"),
		Amount:    1_000_000, // 0.001 SUI
	}

	tx, err := sui_lib.BuildUnsignedTx(params)
	if err != nil {
		t.Fatal(err)
	}

	sig, err := sui_lib.SignTx(params, tx.TxBytes)
	if err != nil {
		t.Fatal(err)
	}

	digest, err := sui_lib.SubmitTx(params, tx, sig)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("tx digest: %s", digest)
}
