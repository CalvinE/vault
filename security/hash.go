package security

import "golang.org/x/crypto/bcrypt"

// HashAndSalt is used to hash and salt a string
func HashAndSalt(s string) (string, error) {
	bytes := []byte(s)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePasswords takes a hashed password and compares it to a plain text password
func ComparePasswords(hashedPassword string, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return false
	}
	return true
}
