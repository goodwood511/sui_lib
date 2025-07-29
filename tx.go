package sui_lib

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/pattonkan/sui-go/sui"
	"github.com/pattonkan/sui-go/suiclient"
	"github.com/pattonkan/sui-go/suisigner"
)

func TransferSUI(params TxParams) (string, error) {
	client := suiclient.NewClient(params.RPCURL)
	seed, err := hex.DecodeString(params.PkHex)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	signer := suisigner.NewSignerByIndex(seed, suisigner.KeySchemeFlagEd25519, 0)
	limit := uint(3)
	CoinType := sui.ObjectTypeFromString(SuiCoinType)
	coinPages, err := client.GetCoins(context.Background(), &suiclient.GetCoinsRequest{
		Owner:    signer.Address,
		CoinType: &CoinType,
		Limit:    limit,
	})

	if err != nil {
		return "", err
	}

	if len(coinPages.Data) == 0 {
		return "", errors.New("no coins found")
	}
	transferCoin := coinPages.Data[0]

	recipient, err := sui.AddressFromHex(params.Recipient)
	if err != nil {
		return "", err
	}

	txn, err := client.TransferSui(
		context.Background(),
		&suiclient.TransferSuiRequest{
			Signer:    signer.Address,
			Recipient: recipient,
			ObjectId:  transferCoin.CoinObjectId,
			Amount:    sui.NewBigInt(params.Amount),
			GasBudget: sui.NewBigInt(suiclient.DefaultGasBudget),
		},
	)
	if err != nil {
		return "", err
	}

	txBytes := txn.TxBytes

	signature, err := signer.SignDigest(txBytes, suisigner.DefaultIntent())
	if err != nil {
		return "", err
	}
	options := &suiclient.SuiTransactionBlockResponseOptions{ShowEffects: true}

	resp, err := client.ExecuteTransactionBlock(
		context.TODO(),
		&suiclient.ExecuteTransactionBlockRequest{
			TxDataBytes: txBytes,
			Signatures:  []*suisigner.Signature{&signature},
			Options:     options,
			RequestType: suiclient.TxnRequestTypeWaitForLocalExecution,
		},
	)
	if err != nil {
		return "", err
	}
	if resp.Effects.Data.IsSuccess() {
		return resp.Digest.String(), nil
	}
	return "", errors.New(resp.Effects.Data.V1.Status.Error)
}
