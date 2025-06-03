package main

import (
	"encoding/json"
	"fmt"
	"go-blockchain-bridge/blockchain"
	"go-blockchain-bridge/ethereum"
	"go-blockchain-bridge/solana"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	ethData := ethereum.FetchEthereumBlock()
	solData := solana.FetchSolanaBalance()
	json.NewEncoder(w).Encode(blockchain.Blockchain)
	if ethData != "" {
		block := blockchain.GenerateBlock(ethData)
		blockchain.Blockchain = append(blockchain.Blockchain, block)

	}
	if solData != "" {
		block := blockchain.GenerateBlock(solData)
		blockchain.Blockchain = append(blockchain.Blockchain, block)
	}
	for _, blk := range blockchain.Blockchain {
		fmt.Fprintf(w, "%+v\n", blk)
	}

}

func main() {
	blockchain.Blockchain = append(blockchain.Blockchain, blockchain.CreateGenesisBlock())
	http.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		handler(w, r)
	})
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
