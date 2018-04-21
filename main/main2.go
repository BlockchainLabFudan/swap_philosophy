package main

import "github.com/scottocs/swap_philosophy/cyb"
import "github.com/scottocs/swap_philosophy/crypto"


func main() {
	initTmpSKForBob()
	initTmpSKForAlice()
	sendBTCToAlice()
	sendCYBToBob()
	go cyb.Run()
	crypto.TestGenerateSharedSecret()
}