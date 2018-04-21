package crypto

import (

	"io/ioutil"
	"strings"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
)

func Readfile(name string)(*btcec.PrivateKey){
	person1,_:=ioutil.ReadFile("./"+name+".txt")
	s:=string(person1)
	lineStr := strings.Split(s,"\n")
	pb,_ := hex.DecodeString(lineStr[0])
	pri1,_:=btcec.PrivKeyFromBytes(btcec.S256(),pb)
	return pri1
}