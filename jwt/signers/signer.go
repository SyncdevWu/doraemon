package signers

type Signer interface {
	Sign(headerBase64, payloadBase64 string) string
	Verify(headerBase64, payloadBase64, signBase64 string) bool
	GetAlgorithm() string
	GetAlgorithmId() string
}
