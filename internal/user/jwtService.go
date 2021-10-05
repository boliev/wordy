package user

import (
	"fmt"
	"time"

	"github.com/boliev/wordy/internal/domain"
	"github.com/dgrijalva/jwt-go"
)

// JwtService struct
type JwtService struct {
	Secret    string
	TokenDays int
}

// CreateJwtService JwtService constructor
func CreateJwtService(secret string, tokenDays int) *JwtService {
	return &JwtService{
		Secret:    secret,
		TokenDays: tokenDays,
	}
}

// Create returns token from user id
func (j JwtService) Create(id int) (*domain.UserAuth, error) {
	expiresAt := time.Now().UTC().AddDate(0, 0, j.TokenDays)
	claims := &jwt.MapClaims{
		"id":        id,
		"expiresAt": expiresAt.Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}

	return &domain.UserAuth{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// Parse returns user id from token
func (j JwtService) Parse(tokenString string) (int, error) {
	claims, err := j.parseClaims(tokenString)

	if err != nil {
		return 0, err
	}

	tm := time.Unix(int64(claims["expiresAt"].(float64)), 0)

	if tm.Before(time.Now()) {
		return 0, fmt.Errorf("token expired")
	}
	id, ok := claims["id"]
	if !ok {
		return 0, fmt.Errorf("wrong token")
	}

	return int(id.(float64)), nil
}

func (j JwtService) parseClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("cant get claims from token")
	}

	if token.Valid != true {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}
