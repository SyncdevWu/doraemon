package jwt

import "time"

const (
	ISSUER     = "iss" // 签发者
	SUBJECT    = "sub" // 面向的用户
	AUDIENCE   = "aud" // 接收方
	EXPIRES_AT = "exp" // 过期时间 必须大于签发时间
	NOT_BEFORE = "nbf" // 生效时间 在这之前jwt是不可用
	ISSUED_AT  = "iat" // 签发时间
	JWT_ID     = "jti" // 唯一身份标识
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
	if p == nil {
		return false
	}
	return true
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
	return p.SetPayload(ISSUER, issuer)
}

func (p *Payload) SetSubject(subject string) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(SUBJECT, subject)
}

func (p *Payload) SetAudience(audience string) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(AUDIENCE, audience)
}

func (p *Payload) SetExpiresAt(expiresAt time.Time) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(EXPIRES_AT, expiresAt)
}

func (p *Payload) SetNotBefore(notBefore time.Time) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(NOT_BEFORE, notBefore)
}

func (p *Payload) SetIssuedAt(issuedAt time.Time) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(ISSUED_AT, issuedAt)
}

func (p *Payload) SetJWTId(jwtId string) *Payload {
	if p.IsNil() {
		return p
	}
	return p.SetPayload(JWT_ID, jwtId)
}
