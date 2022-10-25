package jwt

const (
	ValidationErrorMalformed uint32 = 1 << iota
	ValidationErrorUnverifiable
	ValidationErrorSignatureInvalid
	ValidationErrorAudience
	ValidationErrorExpired
	ValidationErrorIssuedAt
	ValidationErrorIssuer
	ValidationErrorNotValidYet
	ValidationErrorId
	ValidationErrorClaimsInvalid
)
