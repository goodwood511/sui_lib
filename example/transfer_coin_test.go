package example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/goodwood511/sui_lib"
	"github.com/pattonkan/sui-go/sui"
	"github.com/pattonkan/sui-go/suisigner"
	"log"
	"os"
	"testing"

	"github.com/pattonkan/sui-go/suiclient"
)

func TestTransferCoin(t *testing.T) {
	// ====== Configuration ======
	privateKeyHex := os.Getenv("SUI_PRIVATE_KEY_HEX")
	if privateKeyHex == "" {
		t.Skip("❌ Environment variable SUI_PRIVATE_KEY_HEX is not set")
	}
	fmt.Println("🔐 Loaded private key hex")

	sender := os.Getenv("SUI_OWNER_ADDRESS")
	recipientHex := os.Getenv("SUI_RECIPIENT_ADDRESS")
	coinType := os.Getenv("COIN_TYPE")
	amount := sui.NewBigInt(100_000) // 1 token, 6 decimals
	gasBudget := sui.NewBigInt(10000000)

	owner, _ := sui.AddressFromHex(sender)
	recipient, _ := sui.AddressFromHex(recipientHex)

	fmt.Printf("📤 Sender: %s\n", sender)
	fmt.Printf("📥 Recipient: %s\n", recipientHex)
	fmt.Printf("💰 Token Type: %s\n", coinType)
	fmt.Printf("💸 Amount: %s\n", amount.String())

	// === Initialize client ===
	client := suiclient.NewClient(sui_lib.TestnetEndpointUrl)
	fmt.Println("🔌 Connected to Sui mainnet fullnode")

	// === Query token coin object ===
	fmt.Println("🔍 Fetching token coin object...")
	coinObjects, err := client.GetCoins(context.TODO(), &suiclient.GetCoinsRequest{
		Owner:    owner,
		CoinType: &coinType,
		Cursor:   nil,
		Limit:    10,
	})
	if err != nil || len(coinObjects.Data) == 0 {
		log.Fatalf("❌ Failed to find token coin object: %v", err)
	}
	coinID := coinObjects.Data[0].CoinObjectId
	fmt.Println("✅ Token coin object found:", coinID)

	// === Query gas coin (SUI) ===
	suiCoinType := "0x2::sui::SUI"
	fmt.Println("🔍 Fetching gas coin object (SUI)...")
	gasCoins, err := client.GetCoins(context.TODO(), &suiclient.GetCoinsRequest{
		Owner:    owner,
		CoinType: &suiCoinType,
		Cursor:   nil,
		Limit:    10,
	})
	if err != nil || len(gasCoins.Data) == 0 {
		log.Fatalf("❌ Failed to find gas coin object: %v", err)
	}
	gasCoinID := gasCoins.Data[0].CoinObjectId
	fmt.Println("✅ Gas coin object found:", gasCoinID)

	// === Construct transaction (sui_pay) ===
	fmt.Println("🛠️ Building transaction...")
	txn, err := client.Pay(context.TODO(), &suiclient.PayRequest{
		Signer:     owner,
		InputCoins: []*sui.ObjectId{coinID},
		Recipients: []*sui.Address{recipient},
		Amount:     []*sui.BigInt{amount},
		Gas:        gasCoinID,
		GasBudget:  gasBudget,
	})
	if err != nil {
		log.Fatalf("❌ Failed to build transaction: %v", err)
	}
	fmt.Println("✅ Transaction built successfully")
	txBytes := txn.TxBytes

	// === Sign transaction ===
	fmt.Println("✍️ Signing transaction...")
	seed, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatalf("❌ Failed to decode private key: %v", err)
	}
	signer := suisigner.NewSignerByIndex(seed, suisigner.KeySchemeFlagEd25519, 0)

	signature, err := signer.SignDigest(txBytes, suisigner.DefaultIntent())
	if err != nil {
		log.Fatalf("❌ Failed to sign transaction: %v", err)
	}
	fmt.Println("✅ Transaction signed")

	// === Submit transaction ===
	fmt.Println("📡 Submitting transaction to the network...")
	options := &suiclient.SuiTransactionBlockResponseOptions{ShowEffects: true}
	result, err := client.ExecuteTransactionBlock(
		context.TODO(),
		&suiclient.ExecuteTransactionBlockRequest{
			TxDataBytes: txBytes,
			Signatures:  []*suisigner.Signature{&signature},
			Options:     options,
			RequestType: suiclient.TxnRequestTypeWaitForLocalExecution,
		},
	)
	if err != nil {
		log.Fatalf("❌ Failed to submit transaction: %v", err)
	}
	if result.Effects.Data.IsSuccess() {
		fmt.Printf("🎉 Transaction succeeded! TX Digest: %s\n", result.Digest)
	} else {
		fmt.Printf("⚠️ Transaction execution failed. TX Digest: %s\n", result.Digest)
	}
}
