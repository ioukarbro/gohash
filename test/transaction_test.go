package main

import (
	"fmt"
	"gohash/service/tron"
	"gohash/utils/curl"
	"testing"
)

func TestTransaction(T *testing.T) {
	payload := tron.Payload{
		OwnerAddress: tron.OwnerAddress,
		ToAddress:    tron.ToAddress,
		Amount:       1,
	}
}

func TestPostJson(T *testing.T) {
	body, err := curl.PostJson("https://apilist.tronscan.org/api/transfer?address=TYmuaR7B7iPFKFjtZL4L5Jp1hWcfSfbcLB", []byte(""))
	fmt.Println(err)
	fmt.Println(string(body))
}
