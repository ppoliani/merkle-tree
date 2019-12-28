package merkle

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math"
)

func leafHash(leaf []byte) []byte {
	hash := sha256.Sum256(leaf)
	fmt.Printf("%v -> %v \n", HashToString(leaf), HashToString(hash[:]))

	return hash[:]
}

func getSplitPoint(l int) int {
	return int(math.Floor(float64(l) / 2.0))
}

func innerHash(left []byte, right []byte) []byte {
	hash := sha256.Sum256(append(left, right...))
	fmt.Printf("%v | %v -> %v \n", HashToString(left), HashToString(right), HashToString(hash[:]))

	return hash[:]
}

func HashToString(hash []byte) string {
	return base64.URLEncoding.EncodeToString(hash)
}

func CalcMerkleRoot(items [][]byte) []byte {
	switch len(items) {
	case 1:
		return leafHash(items[0])
	default:
		k := getSplitPoint(len(items))
		left := CalcMerkleRoot(items[:k])
		right := CalcMerkleRoot(items[k:])

		return innerHash(left, right)
	}
}
