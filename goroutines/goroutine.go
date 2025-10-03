package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sendSms(s string, ch chan string) {
	time.Sleep(time.Millisecond * 500)
	ch <- fmt.Sprintf("Message send to %s", s)

}

func main() {

	students := []string{"joel", "raj", "Shabana", "Fathima"}
	ch := make(chan string)
	for _, s := range students {
		go sendSms(s, ch)
	}

	msg := <-ch

	fmt.Println(msg)
	wg.Wait()
	// for i := 0; i < len(students); i++ {
	// 	select {
	// 	case msg := <-ch:
	// 		fmt.Println(msg)
	// 	}
	// }
}
