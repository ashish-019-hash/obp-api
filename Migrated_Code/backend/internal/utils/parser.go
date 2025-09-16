package utils

import (
	"math/big"
)

func ParseAmount(amountStr string) *big.Float {
	amount, ok := new(big.Float).SetString(amountStr)
	if !ok {
		return big.NewFloat(0)
	}
	return amount
}

func Contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
