package main

type subscribeReq struct{
	address string  `json:"address" query:"address"`
}
