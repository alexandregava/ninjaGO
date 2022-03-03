package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	senha := "as134679"

	senhaEnc, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(senhaEnc))
}
