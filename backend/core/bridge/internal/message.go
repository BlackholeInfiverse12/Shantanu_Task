package internal

import (
    "crypto/sha256"
    "encoding/hex"
    "sync"
    // "time"
    "fmt"
)

type BridgeMessage struct {
    SourceChain string
    TargetChain string
    Token       string
    Amount      int
    Sender      string
    Receiver    string
    TxHash      string
    Timestamp   int64
    Checksum    string
}

// ComputeChecksum calculates a SHA256 hash of the message fields.
func (m *BridgeMessage) ComputeChecksum() string {
    data := fmt.Sprintf("%s|%s|%s|%d|%s|%s|%s|%d",
        m.SourceChain, m.TargetChain, m.Token, m.Amount,
        m.Sender, m.Receiver, m.TxHash, m.Timestamp)
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])
}

// ValidateChecksum checks if the stored checksum matches the computed one.
func (m *BridgeMessage) ValidateChecksum() bool {
    return m.Checksum == m.ComputeChecksum()
}

// BridgeMessageStore stores processed message hashes to prevent replays.
type BridgeMessageStore struct {
    mu      sync.Mutex
    seenMsg map[string]struct{}
}

func NewBridgeMessageStore() *BridgeMessageStore {
    return &BridgeMessageStore{
        seenMsg: make(map[string]struct{}),
    }
}

// AddIfNew returns true if the message is new, false if duplicate/replay.
func (s *BridgeMessageStore) AddIfNew(msg *BridgeMessage) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    hash := msg.ComputeChecksum()
    if _, exists := s.seenMsg[hash]; exists {
        return false
    }
    s.seenMsg[hash] = struct{}{}
    return true
}