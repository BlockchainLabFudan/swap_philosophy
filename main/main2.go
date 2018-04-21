package main

import (
	"github.com/scottocs/swap_philosophy/runCase"
	"github.com/btcsuite/btcutil"
)


func main() {
	GOD := new(runCase.ExampleCase)
	GOD.InitTmpSKForBob()
	GOD.InitTmpSKForAlice()

	GOD.ExchangePubkeys()

	BtcAmount,_ := btcutil.NewAmount(100)
	GOD.BobInitiate(BtcAmount)
	//GOD.AliceAuditTX()

	//GOD.SendBTCToAlice()
	//GOD.SendCYBToBob()
	//go cyb.Run()
}