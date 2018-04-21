package cyb
//
//import (
//	"bytes"
//	"crypto/rand"
//	"crypto/sha256"
//	"fmt"
//	"time"
//
//	"github.com/scottocs/swap_philosophy/crypto"
//)
//const verify = true
//
//const secretSize = 32
//
//func participate() (error){
//
//
//	var secret [secretSize]byte
//	_, err := rand.Read(secret[:])
//	if err != nil {
//		return err
//	}
//	secretHash := crypto.CalcHash(secret[:], sha256.New())
//
//	// locktime after 500,000,000 (Tue Nov  5 00:53:20 1985 UTC) is interpreted
//	// as a unix time rather than a block height.
//	locktime := time.Now().Add(48 * time.Hour).Unix()
//
//	b, err := buildContract(c, &contractArgs{
//		them:       cmd.cp2Addr,
//		amount:     cmd.amount,
//		locktime:   locktime,
//		secretHash: secretHash,
//	})
//	if err != nil {
//		return err
//	}
//
//	refundTxHash := b.refundTx.TxHash()
//	contractFeePerKb := calcFeePerKb(b.contractFee, b.contractTx.SerializeSize())
//	refundFeePerKb := calcFeePerKb(b.refundFee, b.refundTx.SerializeSize())
//
//	fmt.Printf("Secret:      %x\n", secret)
//	fmt.Printf("Secret hash: %x\n\n", secretHash)
//	fmt.Printf("Contract fee: %v (%0.8f BTC/kB)\n", b.contractFee, contractFeePerKb)
//	fmt.Printf("Refund fee:   %v (%0.8f BTC/kB)\n\n", b.refundFee, refundFeePerKb)
//	fmt.Printf("Contract (%v):\n", b.contractP2SH)
//	fmt.Printf("%x\n\n", b.contract)
//	var contractBuf bytes.Buffer
//	contractBuf.Grow(b.contractTx.SerializeSize())
//	b.contractTx.Serialize(&contractBuf)
//	fmt.Printf("Contract transaction (%v):\n", b.contractTxHash)
//	fmt.Printf("%x\n\n", contractBuf.Bytes())
//	var refundBuf bytes.Buffer
//	refundBuf.Grow(b.refundTx.SerializeSize())
//	b.refundTx.Serialize(&refundBuf)
//	fmt.Printf("Refund transaction (%v):\n", &refundTxHash)
//	fmt.Printf("%x\n\n", refundBuf.Bytes())
//
//	return promptPublishTx(c, b.contractTx, "contract")
//}
