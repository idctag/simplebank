package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuxyz"

// random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// random string with length of n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for range n {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// generate random owner name
func RandomOwner() string {
	return RandomString(6)
}

// generate random money amount
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "MNT", "JPY", "GBP", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
