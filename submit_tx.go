package sui_lib

import (
	"context"
	"errors"
	"github.com/pattonkan/sui-go/suiclient"
	"github.com/pattonkan/sui-go/suisigner"
)

func SubmitTx(params TxParams, tx *suiclient.TransactionBytes, sig *suisigner.Signature) (string, error) {
	cli := suiclient.NewClient(params.RPCURL)
	options := &suiclient.SuiTransactionBlockResponseOptions{ShowEffects: true}

	txBytes := tx.TxBytes

	resp, err := cli.ExecuteTransactionBlock(
		context.TODO(),
		&suiclient.ExecuteTransactionBlockRequest{
			TxDataBytes: txBytes,
			Signatures:  []*suisigner.Signature{sig},
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
