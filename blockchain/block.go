package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}


// create a block caculating hash value with data and prevHash
func CreateBlock(data string, prevHash []byte) *Block {

	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}



// Genesis Block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {

	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	Handle(err)
	return res.Bytes()

}

func Deserialize(data []byte) *Block {

	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	return &block

}

func Handle(err error) {
	if err != nil {
		log.Panic()
	}
}