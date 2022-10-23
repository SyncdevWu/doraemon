package signers

const (
	IdNone = "none"
)

var noneSigner = &NoneSigner{}

type NoneSigner struct{}

func NewNoneSigner() *NoneSigner {
	return noneSigner
}

func (s *NoneSigner) Sign(headerBase64, payloadBase64 string) string {
	return ""
}

func (s *NoneSigner) Verify(headerBase64, payloadBase64, signedBase64 string) bool {
	if s == nil || signedBase64 == "" {
		return false
	}
	return true
}

func (s *NoneSigner) GetAlgorithm() string {
	return IdNone
}
