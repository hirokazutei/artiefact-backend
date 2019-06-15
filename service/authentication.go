package service

import (
	"bytes"

	"golang.org/x/crypto/bcrypt"
)

// PepperAndSaltPassward hashes the password with the pepper given
func PepperAndSaltPassward(password, pepper string) (string, error) {
	// Pepper Password
	var pepperedPassword bytes.Buffer
	pepperedPassword.WriteString(password)
	pepperedPassword.WriteString(pepper)

	// Generate Salt & Hash
	hashByte, err := bcrypt.GenerateFromPassword([]byte(pepperedPassword.String()), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashByte), nil
}

// AuthenticatePassword compares the hashed password
func AuthenticatePassword(rawPassword, hashedPassword, pepper string) (bool, error) {
	rawPasswordHashed, err := PepperAndSaltPassward(rawPassword, pepper)
	if err != nil {
		return false, err
	}
	if rawPasswordHashed == rawPassword {
		return true, nil
	}
	return false, nil
}
