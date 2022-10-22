package signers

const ID_NONE = "none"

type NoneSigner struct{}

func (s *NoneSigner) Sign(headerBase64, payloadBase64 string) string {
	return ""
}

func (s *NoneSigner) Verify(headerBase64, payloadBase64, signBase64 string) bool {
	if s == nil || signBase64 == "" {
		return false
	}
	return true
}

func (s *NoneSigner) GetAlgorithm() string {
	return ID_NONE
}

func (s *NoneSigner) GetAlgorithmId() string {
	return GetId(s.GetAlgorithm())
}
