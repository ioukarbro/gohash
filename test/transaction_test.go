package main

import (
	"encoding/json"
	"fmt"
	"gohash/config"
	"gohash/service/tron"
	"gohash/utils/curl"
	string2 "gohash/utils/string"
	"os"
	"testing"
)

func TestQueryTransaction(T *testing.T) {
	config.Load()
	hash := "1f4d92477549631dc075ac11a3a66814c3a676b91eba8d4bdf10cc5512b0295f"
	result, _ := tron.QueryTransaction(config.TronOwnerAddress, hash)
	string2.PrintJson("result: ", result)
}

func TestCreateTransaction(T *testing.T) {
	config.Load()
	trans, err := tron.CreateTransaction(tron.Payload{
		OwnerAddress: config.TronOwnerAddressHex,
		ToAddress:    config.TronToAddressHex,
		Amount:       1000,
	})
	if err != nil {
		fmt.Println(err)
	}
	string2.LogJson("trans: ", trans)
	b, _ := tron.BroadcastTransaction(trans, config.TronPrivateKey)
	string2.LogJson("bjson: ", b)
}

func TestBroadcastTransaction(T *testing.T) {
	config.Load()
	bt, _ := os.ReadFile("../json1.json")
	var trans tron.Transaction
	_ = json.Unmarshal(bt, &trans)
	string2.LogJson("bt:", trans)
	b, _ := tron.BroadcastTransaction(trans, config.TronPrivateKey)
	string2.LogJson("bjson: ", b)
}

func TestPostJson(T *testing.T) {
	body, err := curl.PostJson("https://apilist.tronscan.org/api/transfer?address=TYmuaR7B7iPFKFjtZL4L5Jp1hWcfSfbcLB", []byte(""))
	fmt.Println(err)
	fmt.Println(string(body))
}
