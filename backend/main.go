package main

import (
	"encoding/json"
	"go-blockchain-bridge/core/blockchain"
	"go-blockchain-bridge/core/ethereum"
	"go-blockchain-bridge/core/solana"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Fetch and append new Ethereum block
	ethData := ethereum.FetchEthereumBlock()
	if ethData != "" {
		block := blockchain.GenerateBlock(ethData)
		blockchain.Blockchain = append(blockchain.Blockchain, block)
	}

	// Fetch and append new Solana block
	solData := solana.FetchSolanaBalance()
	if solData != "" {
		block := blockchain.GenerateBlock(solData)
		blockchain.Blockchain = append(blockchain.Blockchain, block)
	}

	// Return full blockchain as JSON
	err := json.NewEncoder(w).Encode(blockchain.Blockchain)
	if err != nil {
		http.Error(w, "Failed to encode blockchain data", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Initialize the blockchain with a genesis block
	blockchain.Blockchain = append(blockchain.Blockchain, blockchain.CreateGenesisBlock())

	http.HandleFunc("/blocks", handler)

	log.Println("Server started at : http://localhost:8080/blocks")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
