package main

import (
	"encoding/json"
	"fmt"
	"gohash/service/tron"
	string2 "gohash/utils/string"
	"os"
	"testing"
)

func TestJson(T *testing.T) {
	b, err := os.ReadFile("../json1.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var t tron.Transaction
	err = json.Unmarshal(b, &t)
	fmt.Println(err)
	string2.LogJson("json", t)
}
