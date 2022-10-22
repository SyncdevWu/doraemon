package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/syncdevwu/doraemon/core/hashmap"
)

type Claims struct {
	*hashmap.HashMap[string, any]
}

func (c *Claims) NewClaims() *Claims {
	return &Claims{
		hashmap.NewHashMap(&hashmap.Options[string, any]{}),
	}
}

func (c *Claims) IsNil() bool {
	if c == nil {
		return true
	}
	return false
}

func (c *Claims) IsNotNil() bool {
	if c == nil {
		return false
	}
	return true
}

func (c *Claims) Put(claim string, value any) {
	if c.IsNil() {
		return
	}
	c.HashMap.Put(claim, value)
}

func (c *Claims) PutAll(m map[string]any) {
	if c.IsNil() {
		return
	}
	c.HashMap.PutAll(m)
}

func (c *Claims) GetClaim(key string) any {
	if c.IsNil() {
		return nil
	}
	claim, _ := c.HashMap.Get(key)
	return claim
}

func (c *Claims) ToJSONString() string {
	if c.IsNil() {
		return ""
	}
	if data, err := json.Marshal(c.HashMap); err != nil {
		return ""
	} else {
		return string(data)
	}
}

func (c *Claims) Parse(tokenPart string) error {
	if c.IsNil() {
		return errors.New("claim is nil")
	}
	decodeString, err := base64.StdEncoding.DecodeString(tokenPart)
	if err != nil {
		return err
	}
	c.Clear()
	err = json.Unmarshal(decodeString, c.HashMap)
	return err
}

func (c *Claims) MarshalJSON() ([]byte, error) {
	if c.IsNil() {
		return nil, errors.New("claim is nil")
	}
	return json.Marshal(c.HashMap)
}

func (c *Claims) UnmarshalJSON(b []byte) error {
	if c.IsNil() {
		return errors.New("claim is nil")
	}
	c.HashMap.Clear()
	return json.Unmarshal(b, c.HashMap)
}
