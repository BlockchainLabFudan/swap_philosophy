package crypto

import (

	"bytes"
	"crypto/elliptic"
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/btcsuite/btcd/btcec"
	"fmt"
	"github.com/btcsuite/btcutil"

	"encoding/hex"
	"io"
	"os"
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
	//var a []byte
	//err2 := ioutil.WriteFile("./prv.txt", a, 0066)

	fmt.Println(hex.EncodeToString(privKey1.Serialize()))
	fmt.Println(hex.EncodeToString(privKey1.PubKey().SerializeUncompressed()))
	f, _:= os.Create("$GOPATH/bob.txt")
	n, _:= io.WriteString(f , hex.EncodeToString(privKey1.Serialize())) //写入文件(字符串)
	f, _ = os.Create("$GOPATH/alice.txt")
	n, _ = io.WriteString(f , hex.EncodeToString(privKey1.PubKey().SerializeUncompressed())) //写入文件(字符串)

	fmt.Println(n)
	signature,_ :=  privKey1.Sign(btcutil.Hash160([]byte("test message")))
	fmt.Println("verified",signature.Verify(btcutil.Hash160([]byte("test message")), privKey1.PubKey()))

	//pb,_ := hex.DecodeString(hex.EncodeToString(privKey1.Serialize()))
	//privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), pb)
	//fmt.Println(hex.EncodeToString(privKey.Serialize()))
	//fmt.Println(hex.EncodeToString(privKey.PubKey().SerializeUncompressed()))

	secret1 := btcec.GenerateSharedSecret(privKey1, privKey2.PubKey())
	secret2 := btcec.GenerateSharedSecret(privKey2, privKey1.PubKey())

	if !bytes.Equal(secret1, secret2) {
		fmt.Println("no")
	}
	fmt.Println("yes")
}