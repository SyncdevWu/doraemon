package jwt

const (
	ALGORITHM    = "alg"
	TYPE         = "typ"
	CONTENT_TYPE = "cty"
	KEY_ID       = "kid"
)

type Header struct {
	Claims
}

func (h *Header) SetKeyId(keyId string) *Header {
	if h == nil {
		return h
	}
	h.Claims.Put(KEY_ID, keyId)
	return h
}

func (h *Header) AddHeaders(headerClaims map[string]any) *Header {
	if h == nil {
		return h
	}
	h.PutAll(headerClaims)
	return h
}
