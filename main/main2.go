package main

import (
	"github.com/scottocs/swap_philosophy/runCase"
	"github.com/btcsuite/btcutil"
	"fmt"
	"github.com/scottocs/swap_philosophy/crypto"
	"encoding/hex"
	"os"
	"github.com/scottocs/swap_philosophy/cyb"
	"time"
)

func main() {
	GOD := new(runCase.ExampleCase)
	GOD.GlobalAlice = crypto.Getprivatekey(os.Getenv("GOPATH")+"/src/github.com/scottocs/swap_philosophy/alice.txt")
	GOD.GlobalBob = crypto.Getprivatekey(os.Getenv("GOPATH")+"/src/github.com/scottocs/swap_philosophy/bob.txt")
	fmt.Println(hex.EncodeToString(GOD.GlobalAlice.Serialize()))
	GOD.InitTmpKForAliceAndBob()
	//GOD.InitTmpSKForBob()
	//GOD.InitTmpSKForAlice()

	GOD.ExchangePubkeys()

	//fmt.Printf("%v\n\n", hex.EncodeToString(GOD.BobTmpK.Private.PubKey().SerializeUncompressed()))
	//fmt.Printf("%v\n\n", hex.EncodeToString(GOD.BobTmpK.Private.Serialize()))
	fmt.Printf("#####################BobInitiate()#########################\n\n")
	BtcAmount,_ := btcutil.NewAmount(100)
	_, Bobcontract, BobcontractTx, _ := GOD.BobInitiate(BtcAmount)

	BobSig,_ := cyb.GetAcctFromName("bob").TmpK.Private.Sign(btcutil.Hash160(Bobcontract))
	BobSigByte := BobSig.Serialize()
	//_ := string(BobSigByte) //BobSigString
	BobSigHash := btcutil.Hash160(BobSigByte)

	fmt.Printf("#####################AliceAuditTX()#############################\n\n")
	GOD.AliceAuditTX(Bobcontract, BobcontractTx)


	//GOD.SendBTCToAlice()
	//GOD.SendCYBToBob()
	go cyb.Run()

	time.Sleep(10*time.Second)
	signature,_ :=  cyb.GetAcctFromName("bob").TmpK.Private.Sign(BobSigHash)
	cyb.OnReceiveHash("bob", string(signature.Serialize()))

}