package jwt

import (
	"fmt"
	"github.com/syncdevwu/doraemon/jwt/signers"
	"testing"
)

var key = signers.Key{
	PublicKey: []byte("qwertyu"),
}

func TestNewJWT(t *testing.T) {
	jwt := NewJWT(&Options{
		Algorithm: signers.HS256,
	})
	jwt.SetKey(key)
	jwt.
		SetClaim("username", "whb").
		SetClaim("age", 10)
	signingStr := jwt.Sign()
	fmt.Println(signingStr)
}

func TestParseJWT(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOjEwLCJ1c2VybmFtZSI6IndoYiJ9.fN3RTbM7_nemQ_7uPzAGaGeph51nl0diAh1clmy1yHU"
	jwt := ParseJWT(token, key)
	payloads := jwt.GetPayloads()
	username, _ := payloads.Get("username")
	age, _ := payloads.Get("age")
	if res, ok := username.(string); !ok {
		t.Errorf("payloads username expect: %s, but got %s", "whb", res)
	}
	if res, ok := age.(float64); !ok {
		t.Errorf("payloads age expect: %f, but got %f", 10.0, res)
	}
}

func TestJWT_Verify(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOjEwLCJ1c2VybmFtZSI6IndoYiJ9.fN3RTbM7_nemQ_7uPzAGaGeph51nl0diAh1clmy1yHU"
	jwt := ParseJWT(token, key)
	if jwt.Verify(key) != true {
		t.Errorf("payloads age expect: %t, but got %t", true, false)
	}
}
