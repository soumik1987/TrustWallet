package main

type Result struct{
	Id        int         `json:"id"`
	JsonRpc   string      `json:"jsonrpc"`
	Data      BlockData   `json:"result"` 
}

type Transaction struct {
	AccessList           []AccessListData `json:"accessList"`
	BlockHash            string        `json:"blockHash"`
	BlockNumber          string        `json:"blockNumber"`
	ChainId              string        `json:"chainId"`
	From                 string        `json:"from"`
	Gas                  string        `json:"gas"`
	GasPrice             string        `json:"gasPrice"`
	Hash                 string        `json:"hash"`
	Input                string        `json:"input"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	Nonce                string        `json:"nonce"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
	To                   string        `json:"to"`
	TransactionIndex     string        `json:"transactionIndex"`
	Type                 string        `json:"type"`
	V                    string        `json:"v"`
	Value                string        `json:"value"`
	YParity              string        `json:"yParity"`
}

type AccessListData struct {
	Address       string   `json:"address"`
	StorageKeys   []string `json:"storageKeys"`
}

type BlockData struct {
	BaseFeePerGas        string        `json:"baseFeePerGas"`
	BlobGasUsed          string        `json:"blobGasUsed"`
	Difficulty           string        `json:"difficulty"`
	ExcessBlobGas        string        `json:"excessBlobGas"`
	ExtraData            string        `json:"extraData"`
	GasLimit             string        `json:"gasLimit"`
	GasUsed              string        `json:"gasUsed"`
	Hash                 string        `json:"hash"`
	LogsBloom            string        `json:"logsBloom"`
	Miner                string        `json:"miner"`
	MixHash              string        `json:"mixHash"`
	Nonce                string        `json:"nonce"`
	Number               string        `json:"number"`
	ParentBeaconBlockRoot string       `json:"parentBeaconBlockRoot"`
	ParentHash           string        `json:"parentHash"`
	ReceiptsRoot         string        `json:"receiptsRoot"`
	Sha3Uncles           string        `json:"sha3Uncles"`
	Size                 string        `json:"size"`
	StateRoot            string        `json:"stateRoot"`
	Timestamp            string        `json:"timestamp"`
	TotalDifficulty      string        `json:"totalDifficulty"`
	Transactions         []Transaction         `json:"transactions"`
}