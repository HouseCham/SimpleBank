package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnñopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	k := len(charset)

	for i := 0; i < n; i++ {
		c := charset[rand.Intn(k - 1)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalance generates random balance
func RandomBalance() int32 {
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "MXN", "CAD"}
	return currencies[RandomInt(0, 3)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(5))
}