package main

import (
    "encoding/json"
    "log"
    "net/http"
    "bridge/internal"
)

func main() {
    relayHandler := internal.NewRelayHandler()
    bridgeRelay := &internal.BridgeRelay{RelayHandler: relayHandler}

    ethListener, err := internal.NewEthListener(bridgeRelay)
    if err != nil {
        log.Fatal(err)
    }
    solanaListener := internal.NewSolanaListener(bridgeRelay)

    go ethListener.Start()
    go solanaListener.Start()

    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
        events := relayHandler.GetTransactions()
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(events)
    })

    log.Println("Starting web server on http://localhost:8083")
    if err := http.ListenAndServe(":8083", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}