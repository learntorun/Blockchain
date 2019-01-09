package main

import (
	"bytes"
	"crypto/sha256"
)

// basic block struct with field of hash, data, previous hash
type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

// Blockchain struct  with only field of array of block pointer as array
type Blockchain struct {
	blocks []*Block
}

// derive hash data from Block data and previous block hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// create a block caculating hash value with data and prevHash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return block
}

// Add a new block to blockchain
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Genesis Block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Initialize Blockchain with Genesis Block
func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}


func main() {

	chain := InitBlockChain()
	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

}
