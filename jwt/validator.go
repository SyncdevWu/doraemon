package jwt

import "time"

func ValidateTime(jwt *JWT) bool {
	if jwt.IsNil() {
		return false
	}
	now := time.Now()
	// 检查生效时间（生效时间不能晚于当前时间）
	if notBeforeTime, ok := jwt.Payload.GetClaim(NotBefore).(time.Time); ok {
		if ValidateNotAfter(notBeforeTime, now) {
			return false
		}
	}
	// 检查失效时间（失效时间不能早于当前时间）
	if expiresAt, ok := jwt.Payload.GetClaim(ExpiresAt).(time.Time); ok {
		if ValidateNotBefore(expiresAt, now) {
			return false
		}
	}
	// 检查签发时间（签发时间不能晚于当前时间）
	if issusedAt, ok := jwt.Payload.GetClaim(IssuedAt).(time.Time); ok {
		if ValidateNotAfter(issusedAt, now) {
			return false
		}
	}
	return true
}

func ValidateNotAfter(check time.Time, now time.Time) bool {
	return check.After(now)
}

func ValidateNotBefore(check time.Time, now time.Time) bool {
	return check.Before(now)
}
