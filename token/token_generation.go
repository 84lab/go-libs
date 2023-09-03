package token

import (
	"crypto/rand"
	"fmt"
)

//nolint:errcheck
func GenerateRandom(length int) string {
	b := make([]byte, length)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
