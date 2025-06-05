package main

import (
    "log"
    "net/http"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/yourusername/eth-event-listener/internal/listener"
)

const (
    infuraURL       = "https://mainnet.infura.io/v3/688f2501b7114913a6b23a029bd43c9d"
    erc20ContractAddr = "0x26D5Bd2dfEDa983ECD6c39899e69DAE6431Dffbb"
)

func main() {
    client, err := ethclient.Dial(infuraURL)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    contractAddress := common.HexToAddress(erc20ContractAddr)
    eventListener, err := listener.NewEventListener(client, contractAddress)
    if err != nil {
        log.Fatalf("Failed to create event listener: %v", err)
    }

    go eventListener.StartListening()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Ethereum Event Listener is running..."))
    })

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}