package main

import (
	"github.com/scottocs/swap_philosophy/cyb"
)

func main() {
	GOD = new(ExampleCase)

	initTmpSKForBob()
	initTmpSKForAlice()
	sendBTCDepositToAlice()
	sendCYBToBob()
	sendBTCToAlice()
	refundDeposit()
	cyb.Run()
}