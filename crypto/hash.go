package crypto

import "hash"

// Calculate the hash of hasher over buf.
func CalcHash(buf []byte, hasher hash.Hash) []byte {
	hasher.Write(buf)
	return hasher.Sum(nil)
}
