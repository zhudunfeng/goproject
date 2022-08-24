package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// 定义挖矿的难度值
const targetBits = 24

//工作量证明
type ProofOfWork struct {
	block *Block
	//这里的 “目标” ，也就是前一节中所描述的必要条件,hash满足前多少位是0
	//先把哈希转换成一个大整数，然后检测它是否小于目标
	target *big.Int
}

//在 NewProofOfWork 函数中，我们将 big.Int 初始化为 1，然后左移 256 - targetBits 位。
//256 是一个 SHA-256 哈希的位数，我们将要使用的是 SHA-256 哈希算法。target（目标） 的 16 进制形式为：
//获取工作量证明
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//将target设为x << n并返回target（左位移运算）
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

//这里的 nonce，就是上面 Hashcash 所提到的计数器，它是一个密码学术语
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

//  PoW 算法的核心
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	// 	在这个循环中，我们做的事情有：
	// 准备数据
	// 用 SHA-256 对数据进行哈希
	// 将哈希转换成一个大整数
	// 将这个大整数与目标进行比较
	for nonce < maxNonce {
		// 准备数据
		data := pow.prepareData(nonce)
		// 用 SHA-256 对数据进行哈希
		hash = sha256.Sum256(data)
		// 将哈希转换成一个大整数
		hashInt.SetBytes(hash[:])

		// 将这个大整数与目标进行比较
		//func (x *Rat) Cmp(y *Rat) int
		//比较x和y的大小。如x < y返回-1；如x > y返回+1；否则返回0。
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			//增加计数器
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

//工作量证明验证
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
