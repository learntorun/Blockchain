package main

import (
	"bytes"
	"crypto/sha256"
)

// basic block structure
type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

// derive hash data from Block data and previous block hash
func (b *Block) deriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func main() {

}
