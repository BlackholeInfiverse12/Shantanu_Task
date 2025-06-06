package main

import (
    "context"
    "log"
    "net/http"

    "github.com/ethereum/go-ethereum/ethclient"
    "eth-event-listener/internal/listener"
)

const (
    infuraURL         = "wss://mainnet.infura.io/ws/v3/688f2501b7114913a6b23a029bd43c9d"
    erc20ContractAddr = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
)

func main() {
    client, err := ethclient.Dial(infuraURL)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    eventListener, err := listener.NewEventListener(client, erc20ContractAddr)
    if err != nil {
        log.Fatalf("Failed to create event listener: %v", err)
    }

    ctx := context.Background()
    go func() {
        if err := eventListener.ListenTransferEvents(ctx); err != nil {
            log.Fatalf("Error listening for events: %v", err)
        }
    }()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Ethereum Event Listener is running..."))
    })

    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
        data, err := eventListener.GetEventsJSON()
        if err != nil {
            http.Error(w, "Failed to get events", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(data)
    })

    log.Println("Starting server on : http://localhost:8081/events")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}