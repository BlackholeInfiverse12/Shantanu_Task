package ethereum

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func FetchEthereumBlock() string {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/688f2501b7114913a6b23a029bd43c9d")
	if err != nil {
		log.Println("ETH error:", err)
		return ""
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println("ETH header error:", err)
		return ""
	}
	return "Ethereum Block: " + header.Number.String()
}
