package example

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/goodwood511/sui_lib"
)

func TestKeys(t *testing.T) {
	priv, pub, addr, err := sui_lib.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	fmt.Println("Private Key:", hex.EncodeToString(priv))
	fmt.Println("Public Key:", hex.EncodeToString(pub))
	fmt.Println("Sui Address:", addr)
}
