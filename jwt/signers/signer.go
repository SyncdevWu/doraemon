package signers

type Signer interface {
	Sign(headerBase64, payloadBase64 string) string
	Verify(headerBase64, payloadBase64, signedBase64 string) bool
	GetAlgorithm() string
}
