package main

import (
	"fmt"
	"pack/auth"
	"pack/user"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("codersgyan", "secret")

	session := auth.GetSession()
	fmt.Println("session = ", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}
	fmt.Println("user = ", user)
	fmt.Println("user.Email, user.Name = ", user.Email, user.Name)

	color.Green("Prints text in cyan.")

}
