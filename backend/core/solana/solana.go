package solana

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
)

func FetchSolanaBalance() string {
	c := client.NewClient("https://api.devnet.solana.com")

	balance, err := c.GetBalance(context.Background(), "H8jsFz5vF1ejc8BG8Urxu3U7k6pX4Ckn9zMFeZhdQ3Xz")
	if err != nil {
		log.Println("Solana error:", err)
		return ""
	}
	return fmt.Sprintf("Solana Balance: %d", balance)
}
