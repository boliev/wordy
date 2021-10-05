package user

type JwtParser interface {
	Parse(tokenString string) (int, error)
}
