package domain

// UserRepository interface
type UserRepository interface {
	Create(user *User)
	Find(id int) (*User, error)
	FindAll() []*User
}
