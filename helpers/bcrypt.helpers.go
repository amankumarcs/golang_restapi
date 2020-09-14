package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func CompareHashAndPassword(hashPassword string, password []byte) bool {
	byteHash := []byte(hashPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, password)

	if err != nil {
		log.Println(err)
		return false
	}

	return true

}
