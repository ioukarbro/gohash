package tron

import (
	"encoding/json"
	"gohash/utils/curl"
)

const (
	TransferURL             = "http://3.225.171.164:8090/wallet/easytransferbyprivate"
	CreateTransactionURL    = "http://3.225.171.164:8090/wallet/createtransaction"
	GetTransactionSignURL   = "http://3.225.171.164:8090/wallet/gettransactionsign"
	BroadCastTransactionURL = "http://3.225.171.164:8090/wallet/broadcasttransaction"
	PrivateKey              = "d8489f67f90bad8695aa385a9106d45174ae63842ab17ad0b3bdd548d1266c75"
	OwnerAddress            = "41fa27c59b08b08a1e7729a209dd1601b8b8990a33"
	ToAddress               = "41ec74ed234e91c13e2b1efd5dcd4af91434e7305c"
	APIKey                  = "3d89eda8-7359-43fe-a144-e6f8e7dca8c7"
)

type Payload struct {
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
	Amount       int64  `json:"amount"`
}

func CreateTransaction(p Payload) (TxID string, err error) {
	var headers = map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBytes, _ := json.Marshal(p)
	body, _ := curl.PostJsonWithHeader(CreateTransactionURL, jsonBytes, headers)
	var trans Transaction
	if err = json.Unmarshal(body, &trans); err != nil {
		return
	}
	return trans.TxID, err
}

func getTransactionSign(transaction Transaction, privateKey string) (trans TransactionWithSign, err error) {
	payload := struct {
		Transaction Transaction
		PrivateKey  string
	}{
		Transaction: transaction,
		PrivateKey:  privateKey,
	}
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return
	}
	bodyByte, err := curl.PostJson(GetTransactionSignURL, payloadByte)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyByte, &trans)
	return
}

func BroadcastTransaction(transaction Transaction) (b Broadcast, err error) {
	trans, err := getTransactionSign(transaction, PrivateKey)
	if err != nil {
		return
	}
	transByte, err := json.Marshal(trans)
	if err != nil {
		return
	}
	bodyByte, err := curl.PostJson(BroadCastTransactionURL, transByte)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyByte, &b)
	return
}

// QueryTransaction 查询交易
func QueryTransaction(address string) (result ResultData, err error) {
	url := "https://apilist.tronscan.org/api/transfer?limit=20&start=0&sort=-timestamp&count=true&address=" + address
	body, err := curl.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &result)
	return
}
