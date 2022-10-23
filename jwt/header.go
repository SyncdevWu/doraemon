package jwt

const (
	Algorithm   = "alg"
	Type        = "typ"
	ContentType = "cty"
	KeyId       = "kid"
)

type Header struct {
	Claims
}

func (h *Header) SetKeyId(keyId string) *Header {
	if h == nil {
		return h
	}
	h.Claims.Put(KeyId, keyId)
	return h
}

func (h *Header) AddHeaders(headerClaims map[string]any) *Header {
	if h == nil {
		return h
	}
	h.PutAll(headerClaims)
	return h
}
