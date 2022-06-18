package encode

import (
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
)

func Base58ToHexV41(encoded string) (string, error) {
	decoded, _, err := base58.CheckDecode(encoded)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%x", "41", decoded), nil
}
