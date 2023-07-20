package dto

import "time"

type JWTDto struct {
	Sub       string
	ExpiresIn time.Time
}
