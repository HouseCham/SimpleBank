package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnñopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64() * (max - min)
}

func RandomInt(min, max int) int {
	return rand.Intn(max - min) + min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalance generates random balance
func RandomBalance() string {
	return strconv.FormatFloat(RandomFloat(0, 1000), 'f', 2, 64)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "MXN", "CAD"}
	return currencies[RandomInt(0, 4)]
}