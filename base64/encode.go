package base64

import (
	"encoding/base64"
	"fmt"
)

func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("failed to decode a string: %w", err)
	}

	return string(data), nil
}
