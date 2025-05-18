package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x",h.Sum(nil))
}

func generateBlock(oldBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = oldBlock.Index+1
	newBlock.Timestamp =  time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash =oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func createGenesisBlock() Block {
	genesis := Block{0, time.Now().String(), "Genesis", "", ""}
	genesis.Hash = calculateHash(genesis)
	return genesis
}

func main() {
	var Blockchain []Block

	genesisBlock := createGenesisBlock()
	Blockchain =append(Blockchain, genesisBlock)

	// add some block
	for i :=1; i<=3; i++ {
		newBlock :=generateBlock(Blockchain[len(Blockchain)-1], fmt.Sprintf("Block #%d data",i))
		Blockchain = append(Blockchain, newBlock)
	}

	for _,block :=range Blockchain {
		fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\n\n",
			block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash)
	}
}