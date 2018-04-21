package main

import (
	"github.com/scottocs/swap_philosophy/cyb"
	"github.com/scottocs/swap_philosophy/runCase"
)

func main() {
	GOD := new(runCase.ExampleCase)

	GOD.InitTmpSKForBob()
	GOD.InitTmpSKForAlice()
	GOD.SendBTCDepositToAlice()
	GOD.SendCYBToBob()
	GOD.SendBTCToAlice()
	GOD.RefundDeposit()
	cyb.Run()
}