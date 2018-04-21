package crypto

import (

	"bytes"
	"crypto/elliptic"
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/btcsuite/btcd/btcec"
	"fmt"
	"github.com/btcsuite/btcutil"
)
type PrivateKey btcec.PrivateKey
func NewPrivateKey(curve elliptic.Curve) (*btcec.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return (*btcec.PrivateKey)(key), nil
}
func TestGenerateSharedSecret() {
	privKey1, err := NewPrivateKey(btcec.S256())
	if err != nil {

		return
	}
	privKey2, err := NewPrivateKey(btcec.S256())
	if err != nil {

		return
	}
	signature,_ :=  privKey1.Sign(btcutil.Hash160([]byte("test message")))
	fmt.Println("verified",signature.Verify(btcutil.Hash160([]byte("test message")), privKey1.PubKey()))

	secret1 := btcec.GenerateSharedSecret(privKey1, privKey2.PubKey())
	secret2 := btcec.GenerateSharedSecret(privKey2, privKey1.PubKey())

	if !bytes.Equal(secret1, secret2) {
		fmt.Println("no")
	}
	fmt.Println("yes")
}