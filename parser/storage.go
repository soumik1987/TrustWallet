package main

import (
	// "fmt"
	"sync"
)

type Storage interface{
	Subscribe(address string) bool
	SaveTransactionList(txn Transaction)
	FetchTransactionList(address string) []Transaction
	GetCurrentBlock() int64
	SaveCurrentBlock(blockNumber int64)
	IsSubscribed(address string) bool
}


type memoryStorage struct{
	mu sync.Mutex

	lastParsedBlock int64
	subscribedAddress map[string]bool
	dataStorage []Transaction
}

func(ms *memoryStorage) GetCurrentBlock() int64{
	return ms.lastParsedBlock
}

func(ms *memoryStorage) SaveCurrentBlock(blockNumber int64){
	ms.lastParsedBlock = blockNumber
}

func(ms *memoryStorage) SaveTransactionList( txn Transaction){
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.dataStorage = append(ms.dataStorage, txn)
}

func(ms *memoryStorage) FetchTransactionList(address string) []Transaction{
	// loop
	var t = []Transaction{}
	for _, v := range ms.dataStorage {
		if v.From==address || v.To==address{
			t= append(t,v)
		}
	}

	return t
}

func(ms *memoryStorage) Subscribe(address string) bool{
	if _,ok := ms.subscribedAddress[address]; ok{
		return false
	}
	ms.subscribedAddress[address] = true
	return true
}

func(ms *memoryStorage) IsSubscribed(address string) bool{
	_,ok := ms.subscribedAddress[address]
	return ok
}

func NewMemoryStorage() *memoryStorage{
	return &memoryStorage{
		subscribedAddress: make(map[string]bool),
	}
}

// type TransactionList struct{
// 	mu sync.Mutex
// 	// what if the transaction is sent to self. in or out?
// 	InboundTrxs  []Transaction
// 	OutboundTrxs []Transaction
// }

// func(st *TransactionList) SaveTransactions(txn Transaction, txnBound string){
// 	// change to channel
// 	st.mu.Lock()
// 	defer st.mu.Unlock()

// 	if txnBound=="inbound"{
// 		st.InboundTrxs = append(st.InboundTrxs, txn)
// 	}else{
// 		st.OutboundTrxs = append(st.OutboundTrxs, txn)
// 	}
// }

// func(st *TransactionList) FetchTransactions() []Transaction{
// 	return append(st.InboundTrxs, st.OutboundTrxs...)
// }

// func NewTransactionList() *TransactionList{
// 	return &TransactionList{
// 		InboundTrxs: []Transaction{},
// 		OutboundTrxs: []Transaction{},
// 	}
// }
