package main

import (
	"fmt"
	"gohash/utils/encode"
	"testing"
)

func TestString(T *testing.T) {
	// Decode an example Base58Check encoded data.
	encoded := "TXXUYWmkEG8C1KgvxoMsYvH5aBdkcmEoLN"
	fmt.Println(encode.Base58ToHexV41(encoded))
}
