package main

import (
	// "fmt"
)

type Parser interface {

	// last parsed block
	// not the last block published in Eth
	GetCurrentBlock() int

	// add address to observer
	// if address is not valid
	// if address is subscribed
	Subscribe(address string) bool

	// list of inbound or outbound transactions for an address
	// what is the address does not exist or subscribed to
	// address might not be valid
	GetTransactions(address string) ([]Transaction, error)
}

type EthereumParser struct{
	repo Storage
}

func(ep *EthereumParser)  GetCurrentBlock() int{
	return int(ep.repo.GetCurrentBlock())
}

func(ep *EthereumParser)  Subscribe(address string) bool{
	return ep.repo.Subscribe(address)
}

func(ep *EthereumParser)  GetTransactions(address string) ([]Transaction, error){
	return ep.repo.FetchTransactionList(address)
}

func NewEthereumParser(repo Storage) *EthereumParser{
	return &EthereumParser{
		repo: repo,
	}
}
