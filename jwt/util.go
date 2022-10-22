package jwt

import "github.com/syncdevwu/doraemon/jwt/signers"

type JWT struct {
	Header  Header
	Payload Payload
	Signer  signers.Signer
	Tokens  []string
}
