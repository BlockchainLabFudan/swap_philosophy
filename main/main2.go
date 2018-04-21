package main

import "github.com/scottocs/swap_philosophy/cyb"


func main() {
	initTmpSKForBob()
	initTmpSKForAlice()
	sendBTCToAlice()
	sendCYBToBob()
	go cyb.Run()
}