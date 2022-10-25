package jwt

import (
	"errors"
	"time"
)

func ValidateTime(jwt *JWT) *ValidationError {
	vErr := new(ValidationError)
	if jwt.IsNil() {
		return vErr
	}
	now := time.Now()
	// 检查生效时间（生效时间不能晚于当前时间）
	if notBeforeTime, ok := jwt.Payload.GetClaim(NotBefore).(time.Time); ok {
		if validateNotAfter(notBeforeTime, now) {
			vErr.Inner = errors.New("Token is not valid yet")
			vErr.Errors |= ValidationErrorNotValidYet
		}
	}
	// 检查失效时间（失效时间不能早于当前时间）
	if expiresAt, ok := jwt.Payload.GetClaim(ExpiresAt).(time.Time); ok {
		if validateNotBefore(expiresAt, now) {
			vErr.Inner = errors.New("Token is expired")
			vErr.Errors |= ValidationErrorExpired
		}
	}
	// 检查签发时间（签发时间不能晚于当前时间）
	if issusedAt, ok := jwt.Payload.GetClaim(IssuedAt).(time.Time); ok {
		if validateNotAfter(issusedAt, now) {
			vErr.Inner = errors.New("Token used before issued")
			vErr.Errors |= ValidationErrorIssuedAt
		}
	}
	if vErr.valid() != true {
		return vErr
	}
	return nil
}

func validateNotAfter(check time.Time, now time.Time) bool {
	return check.After(now)
}

func validateNotBefore(check time.Time, now time.Time) bool {
	return check.Before(now)
}

func NewValidationError(errorText string, errorFlags uint32) *ValidationError {
	return &ValidationError{
		text:   errorText,
		Errors: errorFlags,
	}
}

type ValidationError struct {
	Inner  error
	Errors uint32
	text   string
}

func (e ValidationError) Error() string {
	if e.Inner != nil {
		return e.Inner.Error()
	} else if e.text != "" {
		return e.text
	} else {
		return "Token is invalid"
	}
}

// No errors
func (e *ValidationError) valid() bool {
	return e.Errors == 0
}
