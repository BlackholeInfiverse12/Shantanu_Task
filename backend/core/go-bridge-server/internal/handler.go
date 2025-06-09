package internal

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "fmt"
    "net/http"
    bc "Projects/TTCont/backend/core/bridge/blockchain"
)

var globalBlockchain = bc.NewBlockchain()

type RelayRequest struct {
    Data string `json:"data"`
}

// Define the expected event structure from /events
type Event struct {
    Index      int    `json:"index"`
    Timestamp  int64  `json:"timestamp"`
    Sourcechain string `json:"Sourcechain"`
    TxHash     string `json:"TxHash"`
    Amount     float64 `json:"Amount"`
}

func fetchEvent() (*Event, error) {
    resp, err := http.Get("http://localhost:8083/events")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var events []Event
    if err := json.Unmarshal(body, &events); err != nil {
        return nil, err
    }
    if len(events) == 0 {
        return nil, fmt.Errorf("no events found")
    }
    return &events[0], nil // Return the first event
}
func HandleEthRelay(w http.ResponseWriter, r *http.Request) {
    event, err := fetchEvent()
    if err != nil {
        log.Printf("Failed to fetch event: %v", err)
        http.Error(w, "Failed to fetch event", http.StatusInternalServerError)
        return
    }

    tx := bc.Transaction{
    Hash:   event.TxHash,
    Amount: fmt.Sprintf("%f", event.Amount), // convert float64 to string if needed
}
    globalBlockchain.PushTransaction(tx)

    w.WriteHeader(http.StatusAccepted)
    json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Ethereum relay processed"})
}

func HandleSolRelay(w http.ResponseWriter, r *http.Request) {
    event, err := fetchEvent()
    if err != nil {
        log.Printf("Failed to fetch event: %v", err)
        http.Error(w, "Failed to fetch event", http.StatusInternalServerError)
        return
    }

    tx := bc.Transaction{
    Hash:   event.TxHash,
    Amount: fmt.Sprintf("%f", event.Amount), // convert float64 to string if needed
}
    globalBlockchain.PushTransaction(tx)

    w.WriteHeader(http.StatusAccepted)
    json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Solana relay processed"})
}