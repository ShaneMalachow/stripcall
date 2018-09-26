package main

import (
	"fmt"
	"github.com/shanemalachow/stripcall/internal/security"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("testpass")
	hash, _ := security.HashNewPassword(password)
	fmt.Println(string(hash))
	result := bcrypt.CompareHashAndPassword(hash, password)
	fmt.Println(result)
}
