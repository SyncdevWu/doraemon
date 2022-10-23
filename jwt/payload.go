package jwt

import "time"

const (
	Issuer    = "iss" // 签发者
	Subject   = "sub" // 面向的用户
	Audience  = "aud" // 接收方
	ExpiresAt = "exp" // 过期时间 必须大于签发时间
	NotBefore = "nbf" // 生效时间 在这之前jwt是不可用
	IssuedAt  = "iat" // 签发时间
	JwtId     = "jti" // 唯一身份标识
)

type Payload struct {
	Claims
}

func (p *Payload) IsNil() bool {
	if p == nil {
		return true
	}
	return false
}

func (p *Payload) IsNotNil() bool {
	return !p.IsNil()
}

func (p *Payload) AddPayloads(payloadClaims map[string]any) *Payload {
	if p.IsNil() {
		return p
	}
	p.Claims.PutAll(payloadClaims)
	return p
}

func (p *Payload) SetPayload(key string, val any) *Payload {
	if p.IsNil() {
		return p
	}
	p.Claims.Put(key, val)
	return p
}

func (p *Payload) SetIssuer(issuer string) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(Issuer, issuer)
}

func (p *Payload) SetSubject(subject string) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(Subject, subject)
}

func (p *Payload) SetAudience(audience string) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(Audience, audience)
}

func (p *Payload) SetExpiresAt(expiresAt time.Time) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(ExpiresAt, expiresAt)
}

func (p *Payload) SetNotBefore(notBefore time.Time) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(NotBefore, notBefore)
}

func (p *Payload) SetIssuedAt(issuedAt time.Time) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(IssuedAt, issuedAt)
}

func (p *Payload) SetJWTId(jwtId string) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(JwtId, jwtId)
}
