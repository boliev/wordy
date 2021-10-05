package user

// JwtParser interface
type JwtParser interface {
	Parse(tokenString string) (int, error)
}
