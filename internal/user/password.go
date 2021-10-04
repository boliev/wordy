package user

import "golang.org/x/crypto/bcrypt"

func hashPassword(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), 8)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func checkPassword(hashed string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
}
