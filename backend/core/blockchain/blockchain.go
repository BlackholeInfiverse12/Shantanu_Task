package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

var Blockchain []Block

func AddMessage(data string) {
    // Create a new block (simplified, you may want to use your real logic)
    newBlock := Block{
        Index:     len(Blockchain),
        Timestamp: "now", // Replace with actual timestamp
        Data:      data,
        PrevHash:  "",
        Hash:      "",
        Nonce:     0,
    }
    Blockchain = append(Blockchain, newBlock)
}

func calculateHash(block Block) string {
	data := fmt.Sprintf("%d, %s, %s, %s, %d", block.Index, block.Timestamp, block.Data, block.PrevHash, block.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func proofOfWork(index int, data, prevHash string) (string, int) {
	nonce := 0
	var hash string
	for {
		blockData := fmt.Sprintf("%d, %s, %s. %d", index, data, prevHash, nonce)
		hashBytes := sha256.Sum256([]byte(blockData))
		hash = hex.EncodeToString(hashBytes[:])
		if hash[:4] == "0000" {
			break
		}
		nonce++
	}
	return hash, nonce
}

func GenerateBlock(data string) Block {
	prev := Blockchain[len(Blockchain)-1]
	hash, nonce := proofOfWork(prev.Index+1, data, prev.Hash)
	block := Block{
		Index:     prev.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prev.Hash,
		Hash:      hash,
		Nonce:     nonce,
	}
	return block
}

func CreateGenesisBlock() Block {
	genesis := Block{0, time.Now().String(), "Genesis", "", "", 0}
	genesis.Hash = calculateHash(genesis)
	return genesis
}
