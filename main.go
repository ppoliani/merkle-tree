package main

import (
	"fmt"
	"merkle-tree/merkle"
)

func main() {
	// fmt.Println(shopping.PriceCheck(4343))
	merkeRoot := merkle.CalcMerkleRoot([][]byte{[]byte("pavlos"), []byte("nikos"), []byte("george"), []byte("michalis")})
	fmt.Printf("Merkel root is %v \n", merkle.HashToString(merkeRoot))
}
