package internal

import (
    "testing"
    "time"
)

func TestBridgeMessageChecksum(t *testing.T) {
    msg := &BridgeMessage{
        SourceChain: "Ethereum",
        TargetChain: "Solana",
        Token:       "USDT",
        Amount:      100,
        Sender:      "0xabc",
        Receiver:    "0xdef",
        TxHash:      "0x123",
        Timestamp:   time.Now().Unix(),
    }
    msg.Checksum = msg.ComputeChecksum()
    if !msg.ValidateChecksum() {
        t.Error("Checksum validation failed")
    }
}

func TestBridgeMessageReplay(t *testing.T) {
    store := NewBridgeMessageStore()
    msg := &BridgeMessage{
        SourceChain: "Ethereum",
        TargetChain: "Solana",
        Token:       "USDT",
        Amount:      100,
        Sender:      "0xabc",
        Receiver:    "0xdef",
        TxHash:      "0x123",
        Timestamp:   1234567890,
    }
    msg.Checksum = msg.ComputeChecksum()
    if !store.AddIfNew(msg) {
        t.Error("First message should be new")
    }
    if store.AddIfNew(msg) {
        t.Error("Duplicate message should be rejected")
    }
}