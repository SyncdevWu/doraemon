package signers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/syncdevwu/doraemon/core/hashmap"
	"github.com/syncdevwu/doraemon/core/util"
	"strings"
)

type RSASigner struct {
	Algorithm string
	Key       Key
	Hash      crypto.Hash
}

var rsaAlgs = hashmap.NewHashMap[string, crypto.Hash](&hashmap.Options[string, crypto.Hash]{})

const (
	RS256 = "RS256"
	RS384 = "RS384"
	RS512 = "RS512"
)

func init() {
	rsaAlgs.Put(RS256, crypto.SHA256)
	rsaAlgs.Put(RS384, crypto.SHA384)
	rsaAlgs.Put(RS512, crypto.SHA512)
}

func NewRSASigner(algorithm string, key Key) *RSASigner {
	if hash, exists := rsaAlgs.Get(algorithm); exists {
		return &RSASigner{
			Algorithm: algorithm,
			Key:       key,
			Hash:      hash,
		}
	} else {
		return &RSASigner{
			Algorithm: RS256,
			Key:       key,
			Hash:      crypto.SHA256,
		}
	}
}

func (R *RSASigner) IsNil() bool {
	if R == nil {
		return true
	}
	return false
}

func (R *RSASigner) IsNotNil() bool {
	return !R.IsNil()
}

func (R *RSASigner) Sign(headerBase64, payloadBase64 string) string {
	if R.IsNil() || !R.Hash.Available() || util.IsNilOrEmpty(R.Key.PublicKey) || util.IsNilOrEmpty(R.Key.PrivateKey) {
		return ""
	}
	block, _ := pem.Decode(R.Key.PrivateKey)
	if block == nil {
		return ""
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return ""
	}
	var builder strings.Builder
	builder.WriteString(headerBase64)
	builder.WriteString(".")
	builder.WriteString(payloadBase64)
	strConcat := builder.String()
	hasher := R.Hash.New()
	hasher.Write([]byte(strConcat))
	if str, err := rsa.SignPKCS1v15(rand.Reader, privateKey, R.Hash, hasher.Sum(nil)); err == nil {
		return util.Base64URLEncode(str)
	} else {
		return ""
	}
}

func (R *RSASigner) Verify(headerBase64, payloadBase64, signedBase64 string) bool {
	if R.IsNil() || !R.Hash.Available() || util.IsNilOrEmpty(R.Key.PublicKey) || util.IsNilOrEmpty(R.Key.PrivateKey) {
		return false
	}
	decodeSignedBase64, err := util.Base64URLDecode(signedBase64)
	if err != nil {
		return false
	}
	publicKey, err := x509.ParsePKCS1PublicKey(R.Key.PublicKey)
	if err != nil {
		return false
	}
	var builder strings.Builder
	builder.WriteString(headerBase64)
	builder.WriteString(".")
	builder.WriteString(payloadBase64)
	strConcat := builder.String()
	hasher := R.Hash.New()
	hasher.Write([]byte(strConcat))
	err = rsa.VerifyPKCS1v15(publicKey, R.Hash, hasher.Sum(nil), decodeSignedBase64)
	if err != nil {
		return false
	}
	return true
}

func (R *RSASigner) GetAlgorithm() string {
	if R.IsNil() {
		return ""
	}
	return R.Algorithm
}
