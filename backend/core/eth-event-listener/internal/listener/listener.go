package listener

import (
    "context"
    "encoding/json"
    "log"
    "math/big"
    "strings"
    "sync"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

const (
    erc20ABI = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
)

type TransferEvent struct {
    From  string   `json:"from"`
    To    string   `json:"to"`
    Value *big.Int `json:"value"`
}

type EventListener struct {
    client          *ethclient.Client
    contractAddress common.Address
    mu              sync.Mutex
    Events          []TransferEvent
}

func NewEventListener(client *ethclient.Client, contractAddress string) (*EventListener, error) {
    address := common.HexToAddress(contractAddress)
    return &EventListener{
        client:          client,
        contractAddress: address,
        Events:          make([]TransferEvent, 0),
    }, nil
}

func (el *EventListener) ListenTransferEvents(ctx context.Context) error {
    transferEvent := "Transfer"
    parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
    if err != nil {
        return err
    }

    query := ethereum.FilterQuery{
        Addresses: []common.Address{el.contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := el.client.SubscribeFilterLogs(ctx, query, logs)
    if err != nil {
        return err
    }
    defer sub.Unsubscribe()

    for {
        select {
        case err := <-sub.Err():
            return err
        case vLog := <-logs:
            event := struct {
                From  common.Address
                To    common.Address
                Value *big.Int
            }{}

            err := parsedABI.UnpackIntoInterface(&event, transferEvent, vLog.Data)
            if err != nil {
                log.Printf("Failed to unpack event: %v", err)
                continue
            }

            if len(vLog.Topics) >= 3 {
                event.From = common.HexToAddress(vLog.Topics[1].Hex())
                event.To = common.HexToAddress(vLog.Topics[2].Hex())
            }

            transfer := TransferEvent{
                From:  event.From.Hex(),
                To:    event.To.Hex(),
                Value: event.Value,
            }

            el.mu.Lock()
            el.Events = append(el.Events, transfer)
            el.mu.Unlock()

            log.Printf("Transfer event: From: %s, To: %s, Value: %s", transfer.From, transfer.To, transfer.Value.String())
        }
    }
}

func (el *EventListener) GetEventsJSON() ([]byte, error) {
    el.mu.Lock()
    defer el.mu.Unlock()
    return json.Marshal(el.Events)
}