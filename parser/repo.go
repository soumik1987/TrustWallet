package main

import (
	"sync"
	"strings"
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
	txnRepo []Transaction
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
	ms.txnRepo = append(ms.txnRepo, txn)
}

func(ms *memoryStorage) FetchTransactionList(address string) []Transaction{
	// loop
	addr := strings.ToLower(address)
	var t = []Transaction{}
	for _, v := range ms.txnRepo {
		if v.From==addr || v.To==addr{
			t= append(t,v)
		}
	}

	return t
}

func(ms *memoryStorage) Subscribe(address string) bool{
	adrs := strings.ToLower(address)
	if _,ok := ms.subscribedAddress[adrs]; ok{
		return false
	}
	ms.subscribedAddress[adrs] = true
	return true
}

func(ms *memoryStorage) IsSubscribed(address string) bool{
	v, ok := ms.subscribedAddress[strings.ToLower(address)]
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