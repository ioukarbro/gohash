package tron

import (
	"encoding/json"
	"gohash/utils/curl"
	string2 "gohash/utils/string"
	"log"
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

func CreateTransaction(p Payload) (trans Transaction, err error) {
	var headers = map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBytes, _ := json.Marshal(p)
	body, _ := curl.PostJsonWithHeader(CreateTransactionURL, jsonBytes, headers)
	string2.LogJson("body: ", string(body))
	if err = json.Unmarshal(body, &trans); err != nil {
		return
	}
	return trans, err
}

func getTransactionSign(transaction Transaction, privateKey string) (trans TransactionWithSign, err error) {
	payload := struct {
		Transaction Transaction `json:"transaction"`
		PrivateKey  string      `json:"privateKey"`
	}{
		Transaction: transaction,
		PrivateKey:  privateKey,
	}
	string2.LogJson("sign payload: ", payload)
	signByte, err := json.Marshal(payload)
	if err != nil {
		return
	}
	bodyByte, err := curl.PostJson(GetTransactionSignURL, signByte)
	string2.LogJson("bodyByte: ", string(bodyByte))
	if err != nil {
		log.Println("curl transaction sign err: ", err)
		return
	}
	err = json.Unmarshal(bodyByte, &trans)
	return
}

func BroadcastTransaction(transaction Transaction, priKey string) (b Broadcast, err error) {
	trans, err := getTransactionSign(transaction, priKey)
	string2.LogJson("trans sign: ", trans)
	if err != nil {
		log.Println("getTransactionSign err: ", err)
		return
	}
	transByte, err := json.Marshal(trans)
	if err != nil {
		return
	}
	bodyByte, err := curl.PostJson(BroadCastTransactionURL, transByte)
	string2.LogJson("broadcast error: ", string(bodyByte))
	if err != nil {
		log.Println("broadcast transaction failed", err)
		return
	}
	err = json.Unmarshal(bodyByte, &b)
	return
}

// QueryTransaction 查询交易
func QueryTransaction(address string, transactionHash string) (resData ResultData, err error) {
	url := "https://apilist.tronscan.org/api/transfer?limit=20&start=0&sort=-timestamp&count=true&address=" + address
	var result TransResult
	println(url)
	body, err := curl.Get(url)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	for _, v := range result.Data {
		if v.TransactionHash == transactionHash {
			return v, err
		}
	}
	return
}
