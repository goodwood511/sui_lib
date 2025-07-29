package example

import (
	"fmt"
	"testing"

	"github.com/goodwood511/sui_lib"
)

func TestGetSuiBalance(t *testing.T) {
	params := sui_lib.SuiBalanceParams{
		Address:  "0x94f1a597b4e8f709a396f7f6b1482bdcd65a673d111e49286c527fab7c2d0961",
		CoinType: "0x2::sui::SUI",
	}

	balance, err := sui_lib.GetSuiBalance(params, sui_lib.TestnetEndpointUrl)
	if err != nil {
		t.Fatalf("GetSuiBalance failed: %v", err)
	}

	fmt.Println("TotalBalance:", balance.TotalBalance)
	fmt.Println("CoinObjectCount:", balance.CoinObjectCount)

	if balance.CoinType != "0x2::sui::SUI" {
		t.Errorf("Unexpected coin type: %s", balance.CoinType)
	}
}
