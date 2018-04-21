package runCase

import (

	ccrypto "github.com/scottocs/swap_philosophy/crypto"
	"github.com/scottocs/swap_philosophy/btc"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/btcec"
	"encoding/json"
	"fmt"
	"time"
	"github.com/scottocs/swap_philosophy/cyb"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

//func (this *ExampleCase)InitTmpSKForBob()  {
//	this.BobTmpK = ccrypto.GenerateTmpKeyPair()
//	//b,_ := this.BobTmpK.ToJson()
//	//this.BobTmpKStr = string(b)
//}
func (this *ExampleCase)GetPKOfBob() string {
	return this.BobTmpK.PubStr
}
func (this *ExampleCase)GetSigOfBob(){

}

//func (this *ExampleCase)InitTmpSKForAlice()  {
//	this.AliceTmpK = ccrypto.GenerateTmpKeyPair()
//	//b,_ := this.AliceTmpK.ToJson()
//	//this.AliceTmpKStr = string(b)
//}
//func (this *ExampleCase)GetPKOfAlice() string {
//	return this.AliceTmpK.PubStr
//}
type Account struct {
	Name      string `json:"name"`
	Balance int64 `json:"balance"`
	Starttime    int64 `json:"st"`
	EndTime     int64 `json:"et"`
	OK bool `json:"ok"`
	SecretHash string `json:"sh"`
	SigB string `json:"sigB"`
	TmpK *ccrypto.Keypair
}

var Acct []Account

func Init() {
	var data = `{"name":"bob","balance":10}`
	act := &Account{}
	err := json.Unmarshal([]byte(data), act)
	if err!=nil {fmt.Println(err)}
	Acct=append(Acct, *act)

	data = `{"name":"alice","balance":100}`
	act = &Account{}
	err = json.Unmarshal([]byte(data), act)
	if err!=nil {fmt.Println(err)}
	Acct=append(Acct, *act)


	data = `{"name":"alice_bob","balance":10,"ok":true}`
	act = &Account{}
	err = json.Unmarshal([]byte(data), act)
	act.Starttime = time.Now().UTC().UnixNano()/1000000
	act.EndTime = act.Starttime+ 24*3600*1000
	//act.SecretHash = string(btcutil.Hash160([]byte("test message")))//Bob gen for CYB
	//act.SigB = act.TmpK.Private.Sign(btcutil.Hash160([]byte("test message")))//Bob gen for BTC
	if err!=nil {fmt.Println(err)}
	Acct=append(Acct, *act)



	fmt.Println(Acct,act)
}
func (this *ExampleCase)GetSigOfAlice(){

}

func (this *ExampleCase)SendBTCDepositToAlice()  {

}
func (this *ExampleCase)SendCYBToBob()  {

}
func (this *ExampleCase)SendBTCToAlice() {

}
func (this *ExampleCase)RefundDeposit()  {

}
func (this *ExampleCase)ExchangePubkeys()  {
	//in example, we ignore the exchange process, just share the PKs.,
}
func (this *ExampleCase)InitTmpKForAliceAndBob(){
	cyb.Init()
}
func (this *ExampleCase)BobInitiate(amount btcutil.Amount) (btcutil.Address,[]byte, *wire.MsgTx, *chainhash.Hash) {
	cmd := btc.NewInitiateCmd(this.BobTmpK.Private.Serialize(),this.BobTmpK.Addr,cyb.GetAcctFromName("alice").TmpK.Addr,amount)
	return cmd.Initiate(nil)
}
func (this *ExampleCase)AliceAuditTX(BobContract []byte, BobContractTx *wire.MsgTx)  {
	cmd := btc.NewAuditContractCmd(BobContract, BobContractTx)
	cmd.Audit()
}

type ExampleCase struct{
	Alice *Account
	Bob *Account
	//txAmountCYB int
	//txAmountBTC int

	//AliceTmpKStr string
	BobTmpK *ccrypto.Keypair
	//BobTmpKStr string
	GlobalAlice *btcec.PrivateKey
	GlobalBob  *btcec.PrivateKey

}