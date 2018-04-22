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
	//Bob prepare its sig and h(sig)
	BobSigString := hex.EncodeToString(BobSigByte) //BobSigString
	BobSigHash := hex.EncodeToString(btcutil.Hash160([]byte(BobSigString)))


	fmt.Printf("#####################AliceAuditTX()#############################\n\n")
	GOD.AliceAuditTX(Bobcontract, BobcontractTx)

	// Below will be the code of a node in CYB chain
	fmt.Printf("##################### clock of CYB Chain  is working()#############################\n\n")
	go cyb.Run()

	fmt.Printf("#####################Get Bob's secret, namely h(SigB),it will be included in Alice's TX######################\n\n")
	bob := cyb.GetAcctFromName("bob")
	bob.SecretHash = BobSigHash
	//signature,_ :=  cyb.GetAcctFromName("bob").TmpK.Private.Sign([]byte(BobSigString))
	fmt.Printf("#####################Bob broadcast its sigB to BTC and CYB #############################\n\n")
	go cyb.OnReceiveHash("bob", BobSigString)

	time.Sleep(10*time.Second)
}