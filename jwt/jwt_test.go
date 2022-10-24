package jwt

import (
	"fmt"
	"github.com/syncdevwu/doraemon/jwt/signers"
	"testing"
)

func TestNewJWT(t *testing.T) {
	jwt := NewJWT(&Options{
		Algorithm: signers.HS256,
	})
	jwt.SetKey(signers.Key{
		PublicKey: []byte("qwertyu"),
	})
	jwt.
		SetClaim("username", "whb").
		SetClaim("age", 10)
	signingStr := jwt.Sign()
	fmt.Println(signingStr)
}

func TestParseJWT(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOjEwLCJ1c2VybmFtZSI6IndoYiJ9.fN3RTbM7_nemQ_7uPzAGaGeph51nl0diAh1clmy1yHU"
	jwt := ParseJWT(token)
	payloads := jwt.GetPayloads()
	fmt.Println(payloads.Keys())
}
