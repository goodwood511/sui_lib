package sui_lib

import (
	"encoding/hex"
	"github.com/pattonkan/sui-go/suisigner"
)

func SignTx(params TxParams, txBytes []byte) (*suisigner.Signature, error) {
	seed, err := hex.DecodeString(params.PkHex)
	if err != nil {
		return nil, err
	}
	signer := suisigner.NewSignerByIndex(seed, suisigner.KeySchemeFlagEd25519, 0)

	signature, err := signer.SignDigest(txBytes, suisigner.DefaultIntent())
	if err != nil {
		return nil, err
	}

	return &signature, nil
}
