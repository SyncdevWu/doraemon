package jwt

type Payload struct {
	Claims
}

func (p *Payload) AddPayloads(payloadClaims map[string]any) *Payload {
	if p == nil {
		return p
	}
	p.Claims.PutAll(payloadClaims)
	return p
}

func (p *Payload) SetPayload(key string, val any) *Payload {
	if p == nil {
		return p
	}
	p.Claims.Put(key, val)
	return p
}
