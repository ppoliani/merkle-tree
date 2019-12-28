package merkle

import (
	"bytes"
)

type Proof struct {
	Root  []byte
	Item  []byte
	Index int
	Path  [][]byte
}

func NewProof(root []byte, item []byte, index int, path [][]byte) *Proof {
	return &Proof{
		Root:  root,
		Index: index,
		Path:  path,
	}
}

func (proof *Proof) Verify() bool {
	hash := LeafHash(proof.Item)
	// clone Path
	path := append([][]byte(nil), proof.Path...)
	path = append(path[:proof.Index], hash)
	path = append(path, proof.Path[proof.Index:]...)

	var reduce func(hash []byte, path [][]byte) []byte
	reduce = func(hash []byte, path [][]byte) []byte {
		if len(path) == 0 {
			return hash
		}

		return reduce(InnerHash(hash, path[0]), path[1:])
	}

	rootHash := reduce(path[0], path[1:])

	return bytes.Compare(rootHash, proof.Root) == 0
}
