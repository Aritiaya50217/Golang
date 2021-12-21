package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func main() {
	// password ที่ต้องการเข้ารหัส(hash)
	password := "secret"
	// ignore error for the sake of simplicity
	hash, _ := HashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hash: ", hash)

	// check ว่า password ก่อนและหลังเข้ารหัส(hash) ตรงกันหรือไม่
	match := CheckPasswordHash(password, hash)
	fmt.Println("Match: ", match)

}
