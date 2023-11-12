package example

import (
	"fmt"
	"time"
)

func ping1S(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("ping : ",i)
		time.Sleep(1*time.Second)
	}
	ch <- 10
}

func sendNoti5S(ch chan string) {
	fmt.Println("send noti")
	time.Sleep(5*time.Second)
	fmt.Println("send noti")
	ch <- "done"
}

func main3() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go ping1S(ch1)
	go sendNoti5S(ch2)

	pingVal, notMess := <-ch1, <-ch2
	fmt.Println(pingVal, notMess)
	fmt.Println("=======")
}