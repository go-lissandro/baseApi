package handler

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func NormalizePassword(p string) []byte {
	return []byte(p)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		NormalizePassword(password),
		bcrypt.DefaultCost,
	)
	fmt.Println("erro ao gerar hash,", err)
	return string(bytes), err
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		NormalizePassword(hash),
		NormalizePassword(password),
	)

	fmt.Println("erro ao checar hash", err)

	return err == nil
}
