package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	storage := NewMemoryStorage()
	url := "https://cloudflare-eth.com"

	// These 2 will be differnet microservices
	ethParser := NewEthereumParser(storage)
	ethSubscriber := NewEthSubscriber(storage, url)

	fmt.Println(ethParser)

	// There will be a persistant value for processedBlock to start
	block, err := ethSubscriber.GetLatestBlockNumber()
	if err != nil {
		log.Fatal(err)
	}
	processedBlock:=int64(block)

	// Poll for new blocks every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Get the latest block number
			latestBlockNumber, err := ethSubscriber.GetLatestBlockNumber()
			if err != nil {
				log.Fatal(err)
			}

			// Process new blocks
			for i := int64(processedBlock); i <= int64(latestBlockNumber); i++ {
				go ethSubscriber.FetchBlockData(big.NewInt(i))
			}
			processedBlock = int64(latestBlockNumber)+1
		}
	}
}
