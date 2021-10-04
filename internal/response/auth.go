package response

import "github.com/boliev/wordy/internal/domain"

// Auth response struct
type Auth struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
}

// CreateAuthFromDomain creates auth response from domain auth
func CreateAuthFromDomain(auth *domain.UserAuth) *Auth {
	return &Auth{
		Token:     auth.Token,
		ExpiresAt: auth.ExpiresAt.String(),
	}
}
