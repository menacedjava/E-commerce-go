package helper

import "golang.org/x/crypto/bcrypt"

// Hash Password
func HashPassword(password string) (string, error) {

	hashPass, eror := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if eror != nil {
		return "", eror
	}
	return string(hashPass), nil
}

// Compare Password
func CompareHashedPassword(password, userPassword string) error {

	eror := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if eror != nil {
		return eror
	}
	return nil
}

