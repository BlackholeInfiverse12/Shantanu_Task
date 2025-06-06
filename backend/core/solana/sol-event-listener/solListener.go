package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "sync"
    "time"
    "strings"
    "fmt"
    bin "github.com/gagliardetto/binary"
    "github.com/gagliardetto/solana-go"
    "github.com/gagliardetto/solana-go/programs/token"
    "github.com/gagliardetto/solana-go/rpc"
)

const (
    solanaRPCURL  = "https://api.mainnet-beta.solana.com"
    tokenMintAddr = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" //USDC H8jsFz5vF1ejc8BG8Urxu3U7k6pX4Ckn9zMFeZhdQ3Xz Es9vMFrzaCER5vL9K1b5nYkZ1u1uY1Y1Y1Y1Y1Y1Y1Y1
    pollInterval  = 5 * time.Second // Increased to avoid rate limit
    maxEvents     = 100             // Only keep the latest 100 events
)

type TransferEvent struct {
    Signature string `json:"signature"`
    Slot      uint64 `json:"slot"`
    From      string `json:"from"`
    To        string `json:"to"`
    Amount    uint64 `json:"amount"`
}

type SolanaListener struct {
    client    *rpc.Client
    tokenMint solana.PublicKey
    mu        sync.Mutex
    events    []TransferEvent
    lastSlot  uint64
}

func NewSolanaListener(rpcURL, mintAddr string) (*SolanaListener, error) {
    mint, err := solana.PublicKeyFromBase58(mintAddr)
    if err != nil {
        return nil, err
    }
    return &SolanaListener{
        client:    rpc.New(rpcURL),
        tokenMint: mint,
        events:    make([]TransferEvent, 0),
    }, nil
}

func (sl *SolanaListener) Listen(ctx context.Context) {
    backoff := pollInterval
    for {
        select {
        case <-ctx.Done():
            return
        default:
            err := sl.pollTransfers()
            if err != nil && isRateLimitError(err) {
                log.Printf("Rate limited, backing off for %v...", backoff)
                time.Sleep(backoff)
                if backoff < 60*time.Second {
                    backoff *= 2
                }
            } else {
                backoff = pollInterval
                time.Sleep(pollInterval)
            }
        }
    }
}

func isRateLimitError(err error) bool {
    if err == nil {
        return false
    }
    return strings.Contains(err.Error(), "Too many requests") || strings.Contains(err.Error(), "429")
}

func (sl *SolanaListener) pollTransfers() error {
    log.Println("Polling for new transfers...")
    limit := 20 // Lowered to reduce RPC load
    sigs, err := sl.client.GetSignaturesForAddressWithOpts(
        context.Background(),
        sl.tokenMint,
        &rpc.GetSignaturesForAddressOpts{
            Limit: &limit,
        },
    )
    if err != nil {
        log.Printf("Error fetching signatures: %v", err)
        return err
    }

    found := false
    for i := len(sigs) - 1; i >= 0; i-- { // Process oldest to newest
        sig := sigs[i]
        if sig.Slot <= sl.lastSlot {
            continue
        }

        txResp, err := sl.client.GetTransaction(
            context.Background(),
            sig.Signature,
            &rpc.GetTransactionOpts{Encoding: solana.EncodingBase64},
        )
        if err != nil || txResp == nil || txResp.Transaction == nil {
            continue
        }

        tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(txResp.Transaction.GetBinary()))
        if err != nil {
            log.Printf("Error decoding transaction: %v", err)
            continue
        }

        for _, inst := range tx.Message.Instructions {
            programID, err := tx.ResolveProgramIDIndex(inst.ProgramIDIndex)
            if err != nil || programID != token.ProgramID {
                continue
            }

            accounts, _ := inst.ResolveInstructionAccounts(&tx.Message)
            decodedInst, err := token.DecodeInstruction(accounts, inst.Data)
            if err != nil {
                continue
            }

            if transferInst, ok := decodedInst.Impl.(*token.Transfer); ok {
                if len(accounts) < 2 || transferInst.Amount == nil {
                    continue
                }

                from := accounts[0].PublicKey.String()
                to := accounts[1].PublicKey.String()

                event := TransferEvent{
                    Signature: sig.Signature.String(),
                    Slot:      sig.Slot,
                    From:      from,
                    To:        to,
                    Amount:    *transferInst.Amount,
                }

                sl.mu.Lock()
                sl.events = append(sl.events, event)
                // Keep only the latest maxEvents
                if len(sl.events) > maxEvents {
                    sl.events = sl.events[len(sl.events)-maxEvents:]
                }
                sl.mu.Unlock()

                log.Printf("Transfer: %s -> %s | Amount: %d | Sig: %s | Slot: %d",
                    event.From, event.To, event.Amount, event.Signature, event.Slot)
                found = true
            }
        }

        if sig.Slot > sl.lastSlot {
            sl.lastSlot = sig.Slot
        }
    }
    if !found {
        log.Println("No new transfer events found in this poll.")
    }
    return nil
}
func (sl *SolanaListener) GetEventsJSON() ([]byte, error) {
    sl.mu.Lock()
    defer sl.mu.Unlock()
    return json.Marshal(sl.events)
}

func main() {
    listener, err := NewSolanaListener(solanaRPCURL, tokenMintAddr)
    if err != nil {
        log.Fatalf("Failed to create Solana listener: %v", err)
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    go listener.Listen(ctx)

   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(`<meta http-equiv="refresh" content="5">`))
    events, err := listener.GetEventsJSON()
    if err != nil {
        http.Error(w, "Failed to get events", http.StatusInternalServerError)
        return
    }
    var parsed []TransferEvent
    if err := json.Unmarshal(events, &parsed); err != nil {
        http.Error(w, "Failed to parse events", http.StatusInternalServerError)
        return
    }
    w.Write([]byte("<h2>Solana Transfer Events</h2><table border='1'><tr><th>Signature</th><th>Slot</th><th>From</th><th>To</th><th>Amount</th></tr>"))
    for i := len(parsed) - 1; i >= 0; i-- {
        e := parsed[i]
        w.Write([]byte(
            "<tr><td>" + e.Signature + "</td><td>" +
                fmt.Sprint(e.Slot) + "</td><td>" +
                e.From + "</td><td>" +
                e.To + "</td><td>" +
                fmt.Sprint(e.Amount) + "</td></tr>"))
    }
    w.Write([]byte("</table>"))
})

    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`<meta http-equiv="refresh" content="5">`))
        data, err := listener.GetEventsJSON()
        if err != nil {
            http.Error(w, "Failed to get events", http.StatusInternalServerError)
            return
        }
        w.Write(data)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8082"
    }
    log.Printf("Solana listener running at http://localhost:%s/events", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}