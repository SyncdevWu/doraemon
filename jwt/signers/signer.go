package signers

import (
	"github.com/syncdevwu/doraemon/core/hashmap"
	"github.com/syncdevwu/doraemon/core/util"
)

type Key struct {
	PublicKey  []byte
	PrivateKey []byte
}

type Signer interface {
	Sign(headerBase64, payloadBase64 string) string
	Verify(headerBase64, payloadBase64, signedBase64 string) bool
	GetAlgorithm() string
}

var Algorithms = hashmap.NewHashMap[string, struct{}](&hashmap.Options[string, struct{}]{})
var empty = struct{}{}

func init() {
	Algorithms.Put(HS256, empty)
	Algorithms.Put(HS384, empty)
	Algorithms.Put(HS512, empty)
	Algorithms.Put(HSHA1, empty)
	Algorithms.Put(RS256, empty)
	Algorithms.Put(RS384, empty)
	Algorithms.Put(RS512, empty)
}

func NewSigner(algorithm string, key Key) Signer {
	if util.IsNilOrEmpty(key.PublicKey) && util.IsNilOrEmpty(key.PrivateKey) {
		return NewNoneSigner()
	}
	if util.IsBlank(algorithm) || None == algorithm {
		return NewNoneSigner()
	}
	// HMAC
	if util.IsNilOrEmpty(key.PrivateKey) {
		return NewHmacSigner(algorithm, key)
	} else {
		// RSA
		return NewRSASigner(algorithm, key)
	}
}
