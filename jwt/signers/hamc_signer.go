package signers

import (
	"crypto"
	"crypto/hmac"
	"github.com/syncdevwu/doraemon/core/hashmap"
	"strings"
)

var hmacAlgs = hashmap.NewHashMap[string, crypto.Hash](&hashmap.Options[string, crypto.Hash]{})

func init() {
	hmacAlgs.Put("HS256", crypto.SHA256)
	hmacAlgs.Put("HS384", crypto.SHA384)
	hmacAlgs.Put("HS512", crypto.SHA512)
	hmacAlgs.Put("HSHA1", crypto.SHA1)
}

type HMacSigner struct {
	Algorithm string
	Key       Key
	Hash      crypto.Hash
}

func NewHmacSigner(algorithm string, key Key) *HMacSigner {
	if hash, exists := hmacAlgs.Get(algorithm); exists {
		return &HMacSigner{
			Algorithm: algorithm,
			Key:       key,
			Hash:      hash,
		}
	} else {
		return &HMacSigner{
			Algorithm: "HS256",
			Key:       key,
			Hash:      crypto.SHA256,
		}
	}
}

func (H *HMacSigner) IsNil() bool {
	if H == nil {
		return true
	}
	return false
}

func (H *HMacSigner) IsNotNil() bool {
	return !H.IsNil()
}

func (H *HMacSigner) Sign(headerBase64, payloadBase64 string) string {
	if H.IsNil() || !H.Hash.Available() {
		return ""
	}
	var builder strings.Builder
	builder.WriteString(headerBase64)
	builder.WriteString(".")
	builder.WriteString(payloadBase64)
	strConcat := builder.String()
	hasher := hmac.New(H.Hash.New, H.Key.PublicKey)
	hasher.Write([]byte(strConcat))
	return encode(hasher.Sum(nil))
}

func (H *HMacSigner) Verify(headerBase64, payloadBase64, signedBase64 string) bool {
	if H.IsNil() {
		return false
	}
	outSignedBase64 := H.Sign(headerBase64, payloadBase64)
	if !hmac.Equal([]byte(outSignedBase64), []byte(signedBase64)) {
		return false
	}
	return true
}

func (H *HMacSigner) GetAlgorithm() string {
	if H.IsNil() {
		return ""
	}
	return H.Algorithm
}
