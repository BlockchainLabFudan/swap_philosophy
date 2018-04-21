package main

import (
	"github.com/scottocs/swap_philosophy/runCase"
	"github.com/btcsuite/btcutil"
	"fmt"
)

func main() {
	GOD := new(runCase.ExampleCase)
	GOD.InitTmpSKForBob()
	GOD.InitTmpSKForAlice()

	GOD.ExchangePubkeys()

	//fmt.Printf("%v\n\n", hex.EncodeToString(GOD.BobTmpK.Private.PubKey().SerializeUncompressed()))
	//fmt.Printf("%v\n\n", hex.EncodeToString(GOD.BobTmpK.Private.Serialize()))
	fmt.Printf("#####################BobInitiate()#########################\n\n")
	BtcAmount,_ := btcutil.NewAmount(100)
	_, Bobcontract, BobcontractTx, _ := GOD.BobInitiate(BtcAmount)

	BobSig,_ := GOD.BobTmpK.Private.Sign(btcutil.Hash160(Bobcontract))
	BobSigByte := BobSig.Serialize()
	BobSigString := string(BobSigByte)
	BobSigHash := btcutil.Hash160(BobSigByte)

	fmt.Printf("#####################AliceAuditTX()#############################\n\n")
	GOD.AliceAuditTX(Bobcontract, BobcontractTx)


	//GOD.SendBTCToAlice()
	//GOD.SendCYBToBob()
	//go cyb.Run()

}