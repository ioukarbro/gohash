package tron

import (
	"encoding/json"
	"gohash/config"
	"gohash/utils/curl"
	"gohash/utils/log"
	string2 "gohash/utils/string"
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
	body, _ := curl.PostJsonWithHeader(config.TronCreateTransactionURL, jsonBytes, headers)
	string2.LogJson("body: ", string(body))
	if err = json.Unmarshal(body, &trans); err != nil {
		log.Sugar.Error(err)
		return
	}
	return trans, err
}

func getTransactionSign(transaction Transaction) (trans TransactionWithSign, err error) {
	payload := struct {
		Transaction Transaction `json:"transaction"`
		PrivateKey  string      `json:"privateKey"`
	}{
		Transaction: transaction,
		PrivateKey:  config.TronPrivateKey,
	}
	string2.LogJson("sign payload: ", payload)
	signByte, err := json.Marshal(payload)
	if err != nil {
		return
	}
	bodyByte, err := curl.PostJson(config.TronTransactionSignURL, signByte)
	string2.LogJson("bodyByte: ", string(bodyByte))
	if err != nil {
		log.Sugar.Error(err)
		return
	}
	err = json.Unmarshal(bodyByte, &trans)
	return
}

func BroadcastTransaction(transaction Transaction) (b Broadcast, err error) {
	trans, err := getTransactionSign(transaction)
	string2.LogJson("trans sign: ", trans)
	if err != nil {
		log.Sugar.Error(err)
		return
	}
	transByte, err := json.Marshal(trans)
	if err != nil {
		log.Sugar.Error(err)
		return
	}
	bodyByte, err := curl.PostJson(config.TronBroadcastTransactionURL, transByte)
	string2.LogJson("broadcast error: ", string(bodyByte))
	if err != nil {
		log.Sugar.Error(err)
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
