package crypto

import (
	"github.com/btcsuite/btcd/btcec"
	"fmt"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"encoding/json"
	"encoding/hex"
)

type Keypair struct{
	private *btcec.PrivateKey
	addr *btcutil.AddressPubKeyHash
	PrvStr	string `json:"private key"`
	PubHash	string `json:"public hash"`
	PubStr	string `json:"public key"`
	AddrStr	string `json:"address"`
}

func GenerateTmpKeyPair() (*Keypair) {
	var NP = new(Keypair)
	NP.private,_ = btcec.NewPrivateKey(btcec.S256())

	NP.PrvStr = hex.EncodeToString(NP.private.Serialize())
	NP.PubStr = hex.EncodeToString(NP.private.PubKey().SerializeCompressed())
	NP.PubHash = string(btcutil.Hash160(NP.private.PubKey().SerializeCompressed()))

	NP.addr, _ = btcutil.NewAddressPubKeyHash([]byte(NP.PubHash),
		&chaincfg.MainNetParams)

	NP.AddrStr = NP.addr.String()

	jsonStr,_ := json.MarshalIndent(NP,"","	")

	fmt.Printf("keyppair:: %v  \n", NP)
	fmt.Printf("jsonStr:: %v  \n", jsonStr)
	return NP
}