package internal

import (
    "context"
    "log"
    "github.com/ethereum/go-ethereum/rpc"
)

type EthListener struct {
    client *rpc.Client
    relay  *BridgeRelay
}

func NewEthListener(relay *BridgeRelay) (*EthListener, error) {
client, err := rpc.Dial("wss://mainnet.infura.io/ws/v3/688f2501b7114913a6b23a029bd43c9d")   
 if err != nil {
        return nil, err
    }
    return &EthListener{client: client, relay: relay}, nil
}

func (el *EthListener) Start() {
    pendingTxs := make(chan string)
    sub, err := el.client.EthSubscribe(context.Background(), pendingTxs, "newPendingTransactions")
    if err != nil {
        log.Fatalf("Failed to subscribe to pending transactions: %v", err)
    }
    defer sub.Unsubscribe()

    for {
        select {
        case txHash := <-pendingTxs:
            el.handleTransaction(txHash)
        }
    }
}

func (el *EthListener) handleTransaction(txHash string) {
    var tx EthTransaction
    err := el.client.Call(&tx, "eth_getTransactionByHash", txHash)
    if err != nil {
        log.Printf("Failed to get transaction: %v", err)
        return
    }

    event := TransactionEvent{
        SourceChain: "Ethereum",
        TxHash:      tx.Hash,
        Amount:      tx.Amount,
    }

    el.relay.PushEvent(event)
    log.Printf("Captured ETH transaction: %s", txHash)
}