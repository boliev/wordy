package domain

// UserRepository interface
type UserRepository interface {
	Create(user *User) error
	Find(id int) (*User, error)
	FindAll() []*User
}
