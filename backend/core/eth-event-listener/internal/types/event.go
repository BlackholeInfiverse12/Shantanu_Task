package types

// TransferEvent represents the structure of an ERC-20 Transfer event.
type TransferEvent struct {
    From  string
    To    string
    Value string
}

// EventType constants for different event types.
const (
    TransferEventType = "Transfer"
)