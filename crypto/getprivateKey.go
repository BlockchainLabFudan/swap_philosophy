package crypto

import (

	//"io/ioutil"

	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	"os"
	"bufio"
	"strings"

)

func Getprivatekey(path string)(*btcec.PrivateKey){
	file,_:=os.Open(path)
	reader:=bufio.NewReader(file)
	str,_:=reader.ReadString('\n')
	str=strings.Trim(str,"\n")
	pb,_ := hex.DecodeString(str)
	pri1,_:=btcec.PrivKeyFromBytes(btcec.S256(),pb)
	return pri1
}
