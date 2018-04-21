package cyb

import (

	"fmt"

	"encoding/json"
	"time"

	"strings"
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
	Sigs AcntWithSig `json:"sigs"`
}
type AcntWithSig struct{
	Mysig string `json:"mine"`
	Othersig string `json:"other"`
}
var acct []Account
func Init() {
	var data = `{"name":"bob","balance":10}`
	act := &Account{}
	err := json.Unmarshal([]byte(data), act)
	if err!=nil {fmt.Println(err)}
	acct=append(acct, *act)

	data = `{"name":"alice","balance":100}`
	act = &Account{}
	err = json.Unmarshal([]byte(data), act)
	if err!=nil {fmt.Println(err)}
	acct=append(acct, *act)


	data = `{"name":"alice_bob","balance":10,"ok":true,"sh":"xxx"}`
	act = &Account{}
	err = json.Unmarshal([]byte(data), act)
	act.Starttime = time.Now().UTC().UnixNano()/1000000
	act.EndTime = act.Starttime+ 24*3600*1000
	if err!=nil {fmt.Println(err)}
	acct=append(acct, *act)



	fmt.Println(acct,act)
}

func hash(sig string) string{
	return sig
}

func fromSpecialAcct2Normal(from Account,to Account)  {
	to.Balance+=from.Balance
	from.Balance = 0
	from.OK = false
}
func getAcctFromName(name string) Account {
	for _, v := range acct {
		if v.Name == name{
			return v
		}
	}
	return Account{}
}
func checkBobBTCSig() bool{
	return true
}
func onReceiveHash(name string,sigs AcntWithSig){
	ac := getAcctFromName(name)
	ac.Sigs = sigs
}
func checkSig(v Account) bool{

	if v.EndTime < time.Now().UTC().UnixNano()/1000000{
	// if Bob provide its sig and all nodes verified it both in BTC and CYB chain,then to Bob
		bobName := strings.Split(v.Name,"_")[1]
		//verify it in CYB
		bob := getAcctFromName(bobName)
		if hash(bob.Sigs.Mysig) != v.SecretHash{
			return false
		}
		//verify it in BTC to guarentee Bob is not cheating
		if checkBobBTCSig() == false{
			return false
		}
		//	to Bob
		fromSpecialAcct2Normal(v,bob)

	}else{
	//	to Alice
		aliceName := strings.Split(v.Name,"_")[0]
		fromSpecialAcct2Normal(v,getAcctFromName(aliceName))

	}
	return true
}

func Run()  {
	Init()
	//for true {

		for _, v := range acct{
			if strings.Index(v.Name,"_") > 0 && v.OK == true{
				checkSig(v)
			}
		}

		time.Sleep(1000000)
	//}
}