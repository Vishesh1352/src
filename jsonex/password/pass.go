package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password1234`
	bs, err := bcrypt.GenerateFromPassword([]bytes(s), bcrypt.Mincost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(bs)
}
