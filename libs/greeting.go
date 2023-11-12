package libs

import "fmt"

func Greeting() {
	fmt.Println("Hello deveploper")
}

func Login(username string, password string) bool {
	if username == "admin" && password == "234" {
		return true
	}
	return false
}