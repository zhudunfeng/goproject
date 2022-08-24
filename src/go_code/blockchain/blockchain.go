package main

//有了区块，下面让我们来实现区块链。本质上，区块链就是一个有着特定结构的数据库，是一个有序，
//每一个块都连接到前一个块的链表。也就是说，区块按照插入的顺序进行存储，每个块都与前一个块相连。
//这样的结构，能够让我们快速地获取链上的最新块，并且高效地通过哈希来检索一个块。
//定义区块链结构体
type Blockchain struct {
	//blocks指针数组
	blocks []*Block
}

//区块链结构体添加区块方法
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
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
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
