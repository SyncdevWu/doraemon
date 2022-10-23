package jwt

import (
	"github.com/syncdevwu/doraemon/core/hashmap"
	"github.com/syncdevwu/doraemon/core/util"
	"github.com/syncdevwu/doraemon/jwt/signers"
	"strings"
)

type JWT struct {
	Header  *Header
	Payload *Payload
	Signer  signers.Signer
	Tokens  []string
}

func NewJWT() *JWT {
	return &JWT{
		Header:  &Header{},
		Payload: &Payload{},
	}
}

func ParseJWT(token string) *JWT {
	if util.IsBlank(token) {
		return nil
	}
	tokens := strings.Split(token, ".")
	if len(tokens) != 3 {
		return nil
	}
	decodeHeader, err := util.Base64URLDecode(tokens[0])
	if err != nil {
		return nil
	}
	header := NewHeader()
	if err = header.Parse(string(decodeHeader)); err != nil {
		return nil
	}
	decodePayload, err := util.Base64URLDecode(tokens[1])
	if err != nil {
		return nil
	}
	payload := NewPayload()
	if err = payload.Parse(string(decodePayload)); err != nil {
		return nil
	}
	return &JWT{
		Header:  header,
		Payload: payload,
		Tokens:  tokens,
	}
}

func (jwt *JWT) SetKey(key signers.Key) *JWT {
	if jwt.IsNil() {
		return jwt
	}
	if util.IsNilOrEmpty(key.PrivateKey) && util.IsNilOrEmpty(key.PublicKey) {
		jwt.SetSigner(signers.None, key)
		return jwt
	}
	algorithm, _ := jwt.Header.GetClaim(Algorithm).(string)
	if util.IsBlank(algorithm) || !signers.Algorithms.ContainsKey(algorithm) {
		// 默认HS256算法 且 保证PublicKey不为空
		key.PublicKey = key.PrivateKey
		jwt.SetSigner(signers.HS256, key)
	} else {
		jwt.SetSigner(algorithm, key)
	}
	return jwt
}

func (jwt *JWT) SetSigner(algorithm string, key signers.Key) *JWT {
	if jwt.IsNil() {
		return jwt
	}
	jwt.Signer = signers.NewSigner(algorithm, key)
	return jwt
}

func (jwt *JWT) GetHeaders() hashmap.HashMap[string, any] {
	headers := hashmap.HashMap[string, any]{}
	if jwt.IsNil() {
		return headers
	}
	keys, vals := jwt.Header.Entries()
	if len(keys) != len(vals) {
		return headers
	}
	for i := 0; i < len(keys); i++ {
		headers.Put(keys[i], vals[i])
	}
	return headers
}

func (jwt *JWT) GetPayloads() hashmap.HashMap[string, any] {
	payloads := hashmap.HashMap[string, any]{}
	if jwt.IsNil() {
		return payloads
	}
	keys, vals := jwt.Payload.Entries()
	if len(keys) != len(vals) {
		return payloads
	}
	for i := 0; i < len(keys); i++ {
		payloads.Put(keys[i], vals[i])
	}
	return payloads
}

func (jwt *JWT) Sign() string {
	if jwt.IsNil() || jwt.Signer == nil {
		return ""
	}
	typ, _ := jwt.Header.GetClaim(Type).(string)
	if util.IsBlank(typ) {
		jwt.Header.Put(Type, "JWT")
	}
	alg, _ := jwt.Header.GetClaim(Algorithm).(string)
	if util.IsBlank(alg) {
		jwt.Header.Put(Algorithm, jwt.Signer.GetAlgorithm())
	}
	headerBase64 := util.Base64URLEncode([]byte(jwt.Header.ToJSONString()))
	payloadBase64 := util.Base64URLEncode([]byte(jwt.Payload.ToJSONString()))
	signedBase64 := jwt.Signer.Sign(headerBase64, payloadBase64)
	var builder strings.Builder
	builder.WriteString(headerBase64)
	builder.WriteString(".")
	builder.WriteString(payloadBase64)
	builder.WriteString(".")
	builder.WriteString(signedBase64)
	return builder.String()
}

func (jwt *JWT) Verify() bool {
	if jwt.IsNil() {
		return false
	}
	if jwt.Signer == nil {
		jwt.Signer = signers.NewNoneSigner()
	}
	tokens := jwt.Tokens
	if util.IsNilOrEmpty(tokens) || len(tokens) != 3 {
		return false
	}
	return jwt.Signer.Verify(tokens[0], tokens[1], tokens[2])
}

func (jwt *JWT) Validate() bool {
	if !jwt.Verify() {
		return false
	}
	// 验证是否Token是否有效
	return ValidateTime(jwt)
}

func (jwt *JWT) IsNil() bool {
	if jwt == nil {
		return true
	}
	return false
}

func (jwt *JWT) IsNotNil() bool {
	return !jwt.IsNil()
}
