package math

import (
	"fmt"
	"strconv"
)

// PreciseFloat32 按小数位截取float
func PreciseFloat32(value float32, precise int) float32 {
	format := fmt.Sprintf("%%.%df", precise)
	v, err := strconv.ParseFloat(fmt.Sprintf(format, value), 32)
	if err != nil {
		return 0
	}
	return float32(v)
}

func Str2Float32(amountStr string) (amount float32, err error) {
	var amount64 float64
	if amount64, err = strconv.ParseFloat(amountStr, 32); err != nil {
		return
	} else {
		return float32(amount64), nil
	}
}
