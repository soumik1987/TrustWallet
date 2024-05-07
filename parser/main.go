package main

import (
	"log"
	"math/big"
	"time"
	"os"

	// "github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func main() {
	ch := make(chan Transaction)
	repo := NewMemoryStorage()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("url")

	// These 2 will be differnet microservices
	ethParser := NewEthereumParser(repo)
	ethSubscriber := NewEthSubscriber(repo, url, ch)

	server := NewServer(ethParser)
	go server.Start()

	// There will be a persistant value for processedBlock to start
	block, err := ethSubscriber.GetLatestBlockNumber()
	if err != nil {
		log.Fatal(err)
	}
	processedBlock:=int64(block)

	// Poll for new blocks every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go ethSubscriber.SaveTransactions()

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
