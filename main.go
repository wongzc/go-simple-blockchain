package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index      int
	Timestamp  string
	Data       string
	PrevHash   string
	Hash       string
	Nonce      int // Number only used once
	Difficulty int // determine how many 0 infront
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func generateBlock(oldBlock Block, data string, difficulty int) Block {
	newBlock := Block{
		Index:      oldBlock.Index + 1,
		Timestamp:  time.Now().String(),
		Data:       data,
		PrevHash:   oldBlock.Hash,
		Difficulty: difficulty,
	}
	mineBlock(&newBlock)
	return newBlock
}

func createGenesisBlock(difficulty int) Block {
	genesis := Block{
		Index:      0,
		Timestamp:  time.Now().String(),
		Data:       "Genesis",
		PrevHash:   "",
		Difficulty: difficulty}
	mineBlock(&genesis)
	return genesis
}

func main() {
	var Blockchain []Block
	difficulty := 4

	genesisBlock := createGenesisBlock(difficulty)
	Blockchain = append(Blockchain, genesisBlock)

	// add some block
	for i := 1; i <= 3; i++ {
		newBlock := generateBlock(Blockchain[len(Blockchain)-1], fmt.Sprintf("Block #%d data", i), difficulty)
		Blockchain = append(Blockchain, newBlock)
	}

	for _, block := range Blockchain {
		fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\nNonce: %d\n",
			block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash, block.Nonce)
	}
}

func mineBlock(block *Block) {
	prefix := strings.Repeat("0", block.Difficulty)
	for {
		hash := calculateHash(*block)
		if strings.HasPrefix(hash, prefix) {
			block.Hash = hash
			break
		}
		block.Nonce++
	}
}
