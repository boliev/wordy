package response

import "github.com/boliev/wordy/internal/domain"

// User response struct
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// CreateUserFromDomain creates user response from domain user
func CreateUserFromDomain(user *domain.User) *User {
	return &User{
		ID:    user.ID,
		Email: user.Email,
	}
}
