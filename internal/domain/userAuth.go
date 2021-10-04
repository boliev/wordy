package domain

import "time"

// UserAuth domain struct
type UserAuth struct {
	Token     string
	ExpiresAt time.Time
}
