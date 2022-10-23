package util

import (
	"encoding/base64"
	"strings"
)

func Base64URLEncode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

func Base64URLDecode(str string) ([]byte, error) {
	if l := len(str) % 4; l > 0 {
		str += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(str)
}
