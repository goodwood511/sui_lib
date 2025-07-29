package sui_lib

import (
	"context"
	"errors"
	"github.com/pattonkan/sui-go/sui"
	"github.com/pattonkan/sui-go/suiclient"
)

func BuildUnsignedTx(params TxParams) (*suiclient.TransactionBytes, error) {
	client := suiclient.NewClient(params.RPCURL)

	recipient, err := sui.AddressFromHex(params.Recipient)
	if err != nil {
		return nil, err
	}

	owner, err := sui.AddressFromHex(params.Owner)
	if err != nil {
		return nil, err
	}

	limit := uint(3)
	CoinType := sui.ObjectTypeFromString(SuiCoinType)
	coinPages, err := client.GetCoins(context.Background(), &suiclient.GetCoinsRequest{
		Owner:    owner,
		CoinType: &CoinType,
		Limit:    limit,
	})

	if err != nil {
		return nil, err
	}

	if len(coinPages.Data) == 0 {
		return nil, errors.New("no coins found")
	}
	transferCoin := coinPages.Data[0]

	txn, err := client.TransferSui(
		context.Background(),
		&suiclient.TransferSuiRequest{
			Signer:    owner,
			Recipient: recipient,
			ObjectId:  transferCoin.CoinObjectId,
			Amount:    sui.NewBigInt(params.Amount),
			GasBudget: sui.NewBigInt(suiclient.DefaultGasBudget),
		},
	)
	if err != nil {
		return nil, err
	}
	return txn, nil
}
