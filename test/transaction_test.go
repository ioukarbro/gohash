package main

import (
	"fmt"
	"gohash/utils/curl"
	"testing"
)

func TestTransaction(T *testing.T) {
	body, _ := curl.Get("https://apilist.tronscan.org/api/transfer?address=TYmuaR7B7iPFKFjtZL4L5Jp1hWcfSfbcLB")
	fmt.Println(string(body))
}

func TestPostJson(T *testing.T) {
	body, err := curl.PostJson("https://apilist.tronscan.org/api/transfer?address=TYmuaR7B7iPFKFjtZL4L5Jp1hWcfSfbcLB", "")
	fmt.Println(err)
	fmt.Println(string(body))
}
