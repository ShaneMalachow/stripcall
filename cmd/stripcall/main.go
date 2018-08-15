package main

import (
	"fmt"
	"github.com/shanemalachow/stripcall/internal/user"
)

func main() {
	newUser, _ := user.CreateUser(1, "Shane", "Malachow", user.Armorer)
	fmt.Println(*newUser)
}
