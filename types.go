package sui_lib

const (
	TestnetEndpointUrl = "https://fullnode.testnet.sui.io:443"

	SuiCoinType = "0x2::sui::SUI"
	USDCoinType = "0x2::usdc::USDC"
)

type TxParams struct {
	RPCURL    string
	PkHex     string
	Owner     string
	Recipient string
	Amount    uint64
}
