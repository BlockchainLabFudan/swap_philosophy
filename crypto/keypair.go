package crypto

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"encoding/hex"
	"encoding/json"
)

type Keypair struct{
	Private *btcec.PrivateKey
	Addr *btcutil.AddressPubKeyHash
	PrvStr	string `json:"private key"`
	PubHash	string `json:"public hash"`
	PubStr	string `json:"public key"`
	AddrStr	string `json:"address"`
}

func GenerateTmpKeyPair() (*Keypair) {
	var NP = new(Keypair)
	NP.Private,_ = btcec.NewPrivateKey(btcec.S256())

	NP.PrvStr = hex.EncodeToString(NP.Private.Serialize())
	NP.PubStr = hex.EncodeToString(NP.Private.PubKey().SerializeCompressed())
	NP.PubHash = string(btcutil.Hash160(NP.Private.PubKey().SerializeCompressed()))

	NP.Addr, _ = btcutil.NewAddressPubKeyHash([]byte(NP.PubHash),
		&chaincfg.MainNetParams)

	NP.AddrStr = NP.Addr.String()

	//jsonStr,_ := json.MarshalIndent(NP,"","	")

	//fmt.Printf("keyppair:: %v  \n", NP)
	//fmt.Printf("jsonStr:: %v  \n", jsonStr)
	return NP
}

func (this *Keypair)ToJson() ([]byte, error){
	return json.MarshalIndent(this, "", "	")
}

func (this *Keypair)FromJson(jsonByte []byte) (error){
	return json.Unmarshal(jsonByte, this)
}