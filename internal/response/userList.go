package response

import "github.com/boliev/wordy/internal/domain"

// UsersList response struct
type UsersList struct {
	Data  []*User `json:"data"`
	Count int     `json:"count"`
}

// CreateUsersListFromDomain creates users list response from domain users list
func CreateUsersListFromDomain(users []*domain.User) *UsersList {
	var responseUsers []*User
	for _, u := range users {
		responseUsers = append(responseUsers, CreateUserFromDomain(u))
	}

	return &UsersList{
		Data:  responseUsers,
		Count: len(responseUsers),
	}
}
