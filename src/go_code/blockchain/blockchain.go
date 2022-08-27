package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

//有了区块，下面让我们来实现区块链。本质上，区块链就是一个有着特定结构的数据库，是一个有序，
//每一个块都连接到前一个块的链表。也就是说，区块按照插入的顺序进行存储，每个块都与前一个块相连。
//这样的结构，能够让我们快速地获取链上的最新块，并且高效地通过哈希来检索一个块。
//定义区块链结构体
// type Blockchain struct {
// 	//blocks指针数组
// 	blocks []*Block
// }

type Blockchain struct {
	//存储区块链的 tip
	//初始时 tip 指向创世块（tip 有尾部，尖端的意思，在这里 tip 存储的是最后一个块的哈希）
	tip []byte
	//数据库连接
	db *bolt.DB
}

//区块链结构体添加区块方法
// func (bc *Blockchain) AddBlock(data string) {
// 	prevBlock := bc.blocks[len(bc.blocks)-1]
// 	newBlock := NewBlock(data, prevBlock.Hash)
// 	bc.blocks = append(bc.blocks, newBlock)
// }
// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	//这是 BoltDB 事务的另一个类型（只读）。在这里，我们会从数据库中获取最后一个块的哈希，
	//然后用它来挖出一个新的块的哈希
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

//区块链迭代器（BlockchainIterator）
// BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

/*
每当要对链中的块进行迭代时，我们就会创建一个迭代器，里面存储了当前迭代的块哈希（currentHash）和数据库的连接（db）。
通过 db，迭代器逻辑上被附属到一个区块链上（这里的区块链指的是存储了一个数据库连接的 Blockchain 实例），
并且通过 Blockchain 方法进行创建：

注意，迭代器的初始状态为链中的 tip，因此区块将从尾到头（创世块为头），也就是从最新的到最旧的进行获取。
实际上，选择一个 tip 就是意味着给一条链“投票”。一条链可能有多个分支，最长的那条链会被认为是主分支。
在获得一个 tip （可以是链中的任意一个块）之后，我们就可以重新构造整条链，找到它的长度和需要构建它的工作。
这同样也意味着，一个 tip 也就是区块链的一种标识符。
*/
// Iterator ...
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

//现在，我们可以实现一个函数来创建有创世块的区块链：
/*
	1.打开一个数据库文件
	2.检查文件里面是否已经存储了一个区块链
	3.如果已经存储了一个区块链：
		i.创建一个新的 Blockchain 实例
		ii.设置 Blockchain 实例的 tip 为数据库中存储的最后一个块的哈希
	4.如果没有区块链：
		i.创建创世块
		ii.存储到数据库
		iii.将创世块哈希保存为最后一个块的哈希
		iv.创建一个新的 Blockchain 实例，初始时 tip 指向创世块（tip 有尾部，尖端的意思，在这里 tip 存储的是最后一个块的哈希）
*/
// func NewBlockchain() *Blockchain {
// 	return &Blockchain{[]*Block{NewGenesisBlock()}}
// }

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	var tip []byte
	//这是打开一个 BoltDB 文件的标准做法
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	//在 BoltDB 中，数据库操作通过一个事务（transaction）进行操作。有两种类型的事务：只读（read-only）和读写（read-write）。
	//这里，打开的是一个读写事务（db.Update(...)），因为我们可能会向数据库中添加创世块。
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			//如果不存在，就生成创世块，创建 bucket，并将区块保存到里面，
			//然后更新 l 键以存储链中最后一个块的哈希
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			//如果存在，就从中读取 l 键
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	/*
		3.如果已经存储了一个区块链：
			i.创建一个新的 Blockchain 实例
			ii.设置 Blockchain 实例的 tip 为数据库中存储的最后一个块的哈希
	*/
	bc := Blockchain{tip, db}

	return &bc
}
