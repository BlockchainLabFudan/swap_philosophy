package main

import (
	"github.com/scottocs/swap_philosophy/crypto"
	"github.com/scottocs/swap_philosophy/cyb"
)

func initTmpSKForBob()  {

}
func getPKOfBob(){

}
func getSigOfBob(){

}

func initTmpSKForAlice()  {

}
func getPKOfAlice(){

}
func getSigOfAlice(){

}

func sendBTCDepositToAlice()  {
	
}
func sendCYBToBob()  {

}
func sendBTCToAlice() {

}
func refundDeposit()  {
	
}
func main() {
	crypto.GenerateTmpKeyPair()


	initTmpSKForBob()
	initTmpSKForAlice()
	sendBTCDepositToAlice()
	sendCYBToBob()
	sendBTCToAlice()
	refundDeposit()
	cyb.Run()
}