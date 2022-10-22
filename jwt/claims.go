package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/syncdevwu/doraemon/core/hashmap"
)

type Claims struct {
	*hashmap.MapWrapper[string, any]
}

func (c *Claims) NewClaims() *Claims {
	return &Claims{
		hashmap.NewMapWrapper(&hashmap.Options[string, any]{}),
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
	c.MapWrapper.Put(claim, value)
}

func (c *Claims) PutAll(m map[string]any) {
	if c.IsNil() {
		return
	}
	c.MapWrapper.PutAll(m)
}

func (c *Claims) GetClaim(key string) any {
	if c.IsNil() {
		return nil
	}
	claim, _ := c.MapWrapper.Get(key)
	return claim
}

func (c *Claims) ToJSONString() string {
	if c.IsNil() {
		return ""
	}
	if data, err := json.Marshal(c.MapWrapper); err != nil {
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
	err = json.Unmarshal(decodeString, c.MapWrapper)
	return err
}

func (c *Claims) MarshalJSON() ([]byte, error) {
	if c.IsNil() {
		return nil, errors.New("claim is nil")
	}
	return json.Marshal(c.MapWrapper)
}

func (c *Claims) UnmarshalJSON(b []byte) error {
	if c.IsNil() {
		return errors.New("claim is nil")
	}
	c.MapWrapper.Clear()
	return json.Unmarshal(b, c.MapWrapper)
}
