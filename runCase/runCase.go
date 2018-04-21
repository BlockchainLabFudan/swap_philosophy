package runCase

import (
	"github.com/scottocs/swap_philosophy/cyb"
	ccrypto "github.com/scottocs/swap_philosophy/crypto"
	"github.com/scottocs/swap_philosophy/btc"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func (this *ExampleCase)InitTmpSKForBob()  {
	this.BobTmpK = ccrypto.GenerateTmpKeyPair()
	//b,_ := this.BobTmpK.ToJson()
	//this.BobTmpKStr = string(b)
}
func (this *ExampleCase)GetPKOfBob() string {
	return this.BobTmpK.PubStr
}
func (this *ExampleCase)GetSigOfBob(){

}

func (this *ExampleCase)InitTmpSKForAlice()  {
	this.AliceTmpK = ccrypto.GenerateTmpKeyPair()
	//b,_ := this.AliceTmpK.ToJson()
	//this.AliceTmpKStr = string(b)
}
func (this *ExampleCase)GetPKOfAlice() string {
	return this.AliceTmpK.PubStr
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
func (this *ExampleCase)BobInitiate(amount btcutil.Amount) (btcutil.Address,[]byte, *wire.MsgTx, *chainhash.Hash) {
	cmd := btc.NewInitiateCmd(this.AliceTmpK.Addr,amount)
	return cmd.Initiate(nil)
}
func (this *ExampleCase)AliceAuditTX(BobContract []byte, BobContractTx *wire.MsgTx)  {
	cmd := btc.NewAuditContractCmd(BobContract, BobContractTx)
	cmd.Audit()
}

type ExampleCase struct{
	Alice *cyb.Account
	Bob *cyb.Account
	txAmountCYB int
	txAmountBTC int
	AliceTmpK *ccrypto.Keypair
	//AliceTmpKStr string
	BobTmpK *ccrypto.Keypair
	//BobTmpKStr string
	GlobalAlice *btcec.PrivateKey
	GlobalBob  *btcec.PrivateKey

}