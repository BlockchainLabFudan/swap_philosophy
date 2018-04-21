package main

import (
	"github.com/scottocs/swap_philosophy/cyb"
	ccrypto "github.com/scottocs/swap_philosophy/crypto"
)

type ExampleCase struct{
	Alice *cyb.Account
	Bob *cyb.Account
	txAmountCYB int
	txAmountBTC int
	AliceTmpK *ccrypto.Keypair
	BobTmpK *ccrypto.Keypair
}
var GOD *ExampleCase

func initTmpSKForBob()  {
	GOD.BobTmpK = ccrypto.GenerateTmpKeyPair()
}
func getPKOfBob(){

}
func getSigOfBob(){

}

func initTmpSKForAlice()  {
	GOD.AliceTmpK = ccrypto.GenerateTmpKeyPair()
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
	GOD = new(ExampleCase)

	initTmpSKForBob()
	initTmpSKForAlice()
	sendBTCDepositToAlice()
	sendCYBToBob()
	sendBTCToAlice()
	refundDeposit()
	cyb.Run()
}