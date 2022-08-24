package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

//定义区块结构体
type Block struct {
	//当前时间戳，也就是区块创建的时间
	Timestamp int64
	//区块存储的实际有效信息，也就是交易
	Data []byte
	//前一个块的哈希，即父哈希
	PrevBlockHash []byte
	//当前块的哈希
	Hash []byte
	//这里的 nonce，就是上面 Hashcash 所提到的计数器，它是一个密码学术语
	Nonce int
}

//定义区块的SetHash方法，引用传递
// func (b *Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)

// 	b.Hash = hash[:]
// }

//序列化
// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

//反序列化
// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

//定义新建区块的函数
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	// block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//为了加入一个新的块，我们必须要有一个已有的块，但是，初始状态下，我们的链是空的，一个块都没有！
//所以，在任何一个区块链中，都必须至少有一个块。这个块，也就是链中的第一个块，通常叫做创世块（genesis block）.
// 让我们实现一个函数来创建创世块：
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
