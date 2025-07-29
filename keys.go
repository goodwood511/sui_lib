package sui_lib

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
)

func GenerateKeyPair() (privKey []byte, pubKey []byte, address string, err error) {
	pubKey, privKey, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, "", err
	}

	prefixedPub := append([]byte{0x00}, pubKey...)
	hash := blake2b.Sum256(prefixedPub)
	address = "0x" + hex.EncodeToString(hash[:])

	return privKey, pubKey, address, nil
}
