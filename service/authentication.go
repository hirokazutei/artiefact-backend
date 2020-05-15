package service

import (
	"golang.org/x/crypto/bcrypt"
)

func pepperPassword(password, pepper string) string {
	return password + pepper
}

// PepperAndSaltPassward hashes the password with the pepper given
func PepperAndSaltPassward(password, pepper string) (string, error) {
	// Generate Salt & Hash
	hashByte, err := bcrypt.GenerateFromPassword([]byte(pepperPassword(password, pepper)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashByte), nil
}

// AuthenticatePassword compares the hashed password
func AuthenticatePassword(rawPassword, hashedPassword, pepper string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pepperPassword(rawPassword, pepper)))
	if err == nil {
		return true, nil
	}
	return false, nil
}
