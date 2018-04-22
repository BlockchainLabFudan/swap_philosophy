package cyb

import (

	"fmt"
	ccrypto "github.com/scottocs/swap_philosophy/crypto"
	"encoding/json"
	"time"
	"github.com/btcsuite/btcutil"
	"strings"
	"encoding/hex"
)

//var accnt = `{"account":[{"name":"bob","balance":10},{"name":"alice","balance":100}]}`
//var accnt2 = json.Unmarshal([]byte(`{{"name":"bob","balance":10},{"name":"alice","balance":100},{"name":"alice_bob","balance":10,"st":1524277394,"et":1624277394}]}`))
type Account struct {
	Name      string `json:"name"`
	Balance int64 `json:"balance"`
	Starttime    int64 `json:"st"`
	EndTime     int64 `json:"et"`
	OK bool `json:"ok"`
	SecretHash string `json:"sh"`
	SigB string `json:"sigB"`
	TmpK *ccrypto.Keypair
}

var acct []*Account
func Init() {
	var data = `{"name":"bob","balance":10}`
	act := &Account{}
	act.TmpK = ccrypto.GenerateTmpKeyPair()
	err := json.Unmarshal([]byte(data), act)
	if err!=nil {fmt.Println(err)}
	acct=append(acct, act)

	data = `{"name":"alice","balance":100}`
	act = &Account{}
	act.TmpK = ccrypto.GenerateTmpKeyPair()
	err = json.Unmarshal([]byte(data), act)
	if err!=nil {fmt.Println(err)}
	acct=append(acct, act)


	data = `{"name":"alice_bob","balance":10,"ok":true}`
	act = &Account{}
	err = json.Unmarshal([]byte(data), act)
	act.Starttime = time.Now().UTC().UnixNano()/1000000
	act.EndTime = act.Starttime+ 24*3600*1000
	act.SecretHash = string(btcutil.Hash160([]byte("test message")))//Bob gen for CYB

	//signature,_ := GetAcctFromName("bob").TmpK.Private.Sign(btcutil.Hash160([]byte("test message")))
	//act.SigB = string(signature.Serialize()) //Bob gen for BTC
	if err!=nil {fmt.Println(err)}
	acct=append(acct, act)



	fmt.Println(acct,act)
}

func hash(sig string) string{
	return hex.EncodeToString(btcutil.Hash160(([]byte)(sig)))
}

func fromSpecialAcct2Normal(from *Account,to *Account)  {
	to.Balance+=from.Balance
	from.Balance = 0
	from.OK = false
}
func GetAcctFromName(name string) *Account {
	for _, v := range acct {
		if v.Name == name{

			return v
		}
	}
	return &Account{}
}
func checkBobBTCSig(bob *Account) bool{
	BTCChainBobsig := bob.SigB
	return bob.SigB == BTCChainBobsig
}
func OnReceiveHash(name string,sig string){

	for _, v := range acct {
		if v.Name == name{
			v.SigB = sig
			break
		}
	}
}
func checkSig(v *Account) bool{

	if v.EndTime > time.Now().UTC().UnixNano()/1000000{
	// if Bob provide its sig and all nodes verified it both in BTC and CYB chain,then to Bob
		bobName := strings.Split(v.Name,"_")[1]
		//verify it in CYB
		bob := GetAcctFromName(bobName)
		//fmt.Println(bob.SigB,bob.SecretHash)
		if hash(bob.SigB) != bob.SecretHash{
			fmt.Println("bob.SigB != bob.SecretHash")
			return false
		}
		//verify it in BTC to guarentee Bob is not cheating
		if checkBobBTCSig(bob) == false{
			fmt.Println("bob cheat with different sig in cyb and btc net")
			return false
		}
		fmt.Println("Bob redeem ok,bob from ",bob.Balance,"||", v.Name," from",v.Balance)
		//	to Bob
		fromSpecialAcct2Normal(v,bob)
		fmt.Println("Bob redeem ok,bob to ",bob.Balance,"||",v.Name," to",v.Balance)

	}else{
	//	to Alice
		aliceName := strings.Split(v.Name,"_")[0]
		fromSpecialAcct2Normal(v,GetAcctFromName(aliceName))
		fmt.Println("refund to alice")
	}
	return true
}

func Run()  {

	for true {

		for _, v := range acct{
			if strings.Index(v.Name,"_") > 0 && v.OK == true{
				checkSig(v)
			}
		}

		time.Sleep(time.Second)
	}
}