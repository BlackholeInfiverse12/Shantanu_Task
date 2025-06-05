package listener

import (
    "context"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

const (
    erc20ABI = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
)

type EventListener struct {
    client      *ethclient.Client
    contractAddress common.Address
}

func NewEventListener(client *ethclient.Client, contractAddress string) (*EventListener, error) {
    address := common.HexToAddress(contractAddress)
    return &EventListener{
        client:         client,
        contractAddress: address,
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

    logs := make(chan ethereum.Log)
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

            log.Printf("Transfer event: From: %s, To: %s, Value: %s", event.From.Hex(), event.To.Hex(), event.Value.String())
        }
    }
}