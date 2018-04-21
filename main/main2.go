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

	fmt.Printf("#####################BobInitiate()#########################\n\n")
	BtcAmount,_ := btcutil.NewAmount(100)
	_, Bobcontract, BobcontractTx, _ := GOD.BobInitiate(BtcAmount)

	fmt.Printf("#####################AliceAuditTX()#############################\n\n")
	GOD.AliceAuditTX(Bobcontract, BobcontractTx)


	//GOD.SendBTCToAlice()
	//GOD.SendCYBToBob()
	//go cyb.Run()

}