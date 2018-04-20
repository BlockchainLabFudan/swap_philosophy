package node

import (
	"github.com/tidwall/gjson"
	"fmt"
	"time"
	"strings"
)

var accnt = `{"account":[{"name":"bob","balance":10},{"name":"alice","balance":100}]}`

func checkSig(){

}

func Run()  {
	for true {

		arr := gjson.Get(accnt, "account")
		re:= arr.Array()
		for _, v := range re {
			name := gjson.Get(v.String(), "name").String()
			if strings.Index(name,"_") > 0{
				checkSig()
			}

		}

		time.Sleep(1)
	}
}