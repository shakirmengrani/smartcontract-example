package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type block struct {
	index     int
	timestamp int64
	data      []byte
	prevHash  []byte
	hash      []byte
}

func (b *block) calculateHash() {
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	headers := bytes.Join([][]byte{b.prevHash, b.data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.hash = hash[:]
}

type blockchain struct {
	chain []*block
}

func (bc *blockchain) init() {
	newBlock := block{index: 0, timestamp: time.Now().Unix(), prevHash: []byte(""), data: []byte("New chain")}
	newBlock.calculateHash()
	bc.chain = append(bc.chain, &newBlock)
}

func (bc *blockchain) addBlock(newBlock *block) {
	newBlock.prevHash = bc.chain[len(bc.chain)-1].hash
	newBlock.calculateHash()
	bc.chain = append(bc.chain, newBlock)
}

func main() {
	coin := blockchain{}
	coin.init()
	for i := 1; i <= 10; i++ {
		coin.addBlock(&block{index: 1, timestamp: time.Now().Unix(), data: []byte(fmt.Sprintf("amount %d", 8*i))})
	}

	for _, val := range coin.chain {
		fmt.Println("prevHash: " + string(val.prevHash))
		fmt.Println("Index: " + string(val.index))
		fmt.Println("Data: " + string(val.data))
		fmt.Println("Timestamp: " + string(val.timestamp))
		fmt.Println("hash: " + string(val.hash))
		fmt.Println("--------------------------------------------------------")
	}

}
