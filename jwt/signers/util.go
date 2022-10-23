package signers

import (
	"encoding/base64"
	"github.com/syncdevwu/doraemon/core/util"
	"strings"
)

type Key struct {
	PublicKey  []byte
	PrivateKey []byte
}

func None() *NoneSigner {
	return NewNoneSigner()
}

func HS256(key []byte) {

}

func encode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

func decode(str string) ([]byte, error) {
	if l := len(str) % 4; l > 0 {
		str += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(str)
}

func createSigner(algorithm string, key Key) Signer {
	if util.IsNilOrEmpty(key.PublicKey) && util.IsNilOrEmpty(key.PrivateKey) {
		return None()
	}
	if util.IsBlank(algorithm) || IdNone == algorithm {
		return None()
	}
	// 对称加密
	if util.IsNilOrEmpty(key.PrivateKey) {
		return NewHmacSigner(algorithm, key)
	} else {
		return NewRSASigner(algorithm, key)
	}
}
