package main

import (
	"fmt"
	"log"
	"math/big"

	"bytes"
	"encoding/json"
	"net/http"
)

type Subscriber interface {
	// Subscribe(address string) bool
	FetchBlockData(blockNumber *big.Int)
	GetLatestBlockNumber() (uint64, error)
}

type EthSubscriber struct{
	storage Storage
	url string
}

func NewEthSubscriber(storage Storage, url string) *EthSubscriber{
	return &EthSubscriber{
		storage: storage,
		url: url,
	}
}

func(es *EthSubscriber) FetchBlockData(blockNumber *big.Int){
	blockNumberHex := fmt.Sprintf("0x%x", blockNumber)
	payload := map[string]interface{}{
		"method": "eth_getBlockByNumber",
		"params": []interface{}{
			blockNumberHex,
			true,
		},
		"id":      1,
		"jsonrpc": "2.0",
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	// Send POST request
	resp, err := http.Post(es.url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read response body
	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("Error Decoding: ", resp.Body, err)
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error marshaling map to JSON:", err)
		return
	}

	var block Result
	err = json.Unmarshal([]byte(jsonData), &block)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return
	}

	for _, v := range block.Data.Transactions {
		go es.processTransaction(v)
	}
}


func(es *EthSubscriber) GetLatestBlockNumber() (uint64, error) {
	payload := map[string]interface{}{
		"method": "eth_blockNumber",
		"params": []interface{}{},
		"id":      2,
		"jsonrpc": "2.0",
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	// Send POST request
	resp, err := http.Post(es.url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}
	
	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	latestBlockNumber := new(big.Int)
	latestBlockNumber, ok := latestBlockNumber.SetString(result["result"].(string)[2:], 16)
	if !ok {
		return 0, fmt.Errorf("failed to convert block number")
	}

	es.storage.SaveCurrentBlock(latestBlockNumber.Int64())

	return latestBlockNumber.Uint64(), nil
}


func(es *EthSubscriber) processTransaction(transaction Transaction) {
	// handle txn without To field
	fmt.Printf("Hash: %s\n", transaction.Hash)
	fmt.Printf("From: %s\n", transaction.From)
	fmt.Printf("To: %s\n", transaction.To)
	if es.storage.IsSubscribed(transaction.From) || ( transaction.To!="" && es.storage.IsSubscribed(transaction.To) ){
		// fmt.Printf("Transaction Hash: %s\n", transaction.Hash )
		fmt.Printf("From: %v\n", transaction)
		es.storage.SaveTransactionList(transaction)
	}

	fmt.Println("-------------------------")
}
