package example

import (
	"fmt"
	"time"
)

func isInSystem(username string) bool {
	return true
}

func getUserDetail(username string) (int, string) {
	return 201, "manager"
}

func getDaparture(username string, departure *string) {
	if username != "" {
		*departure = "home"
	}
} 

func checkLogin(username string, password string) {
	if isInSystem(username) {
		fmt.Println("found user in system")
		getUserDetail(username)
		departure := ""
		getDaparture(username, &departure)
	}
}

func logEnd() {
	time.Now()
	fmt.Println("completed program")
	fmt.Println(time.Now())
}

func getMember() {
	fmt.Println("Please Wait...")
	time.Sleep(3*time.Second)
}

func checkServerRespone() {
	fmt.Println("check server time")
	time.Sleep(3*time.Second)
	panic("server error")
}

// main
func examFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover ")
			fmt.Println(r)
		}
	}()
	defer logEnd()
	getMember()
	checkServerRespone()
	
}