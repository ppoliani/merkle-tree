package main

import (
	"fmt"
	"merkle-tree/merkle"
)

func main() {
	// fmt.Println(shopping.PriceCheck(4343))
	merkeRoot := merkle.CalcMerkleRoot([][]byte{[]byte("pavlos"), []byte("nikos"), []byte("george"), []byte("michalis")})
	fmt.Printf("Merkel root is %v \n", merkle.HashToString(merkeRoot))

	// proof
	item := "nikos"
	proof := merkle.Proof {
		Root: merkeRoot,
		Index: 1,
		Item: []byte(item),
		Path: [][]byte{
			merkle.LeafHash([]byte("pavloss")), 
			merkle.InnerHash(
				merkle.LeafHash([]byte("george")), 
				merkle.LeafHash([]byte("michalis")),
			),
		},
	}
	
	fmt.Printf("Value %v exists in the tree %v \n", item, proof.Verify())
}
