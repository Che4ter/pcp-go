package main

import (
	"fmt"
	"time"
)

func talkingback(receiver string, msg chan string) {
	for i := 0; ; i++ {
		msg <- fmt.Sprintf("Hello %s : %d", receiver, i)
		<-time.After(time.Millisecond * 200)
	}
}

func main() {
	msg := make(chan string)
	fmt.Println("Start talking")
	go talkingback("Alan", msg)
	go talkingback("Philipp", msg)
	for i := 0; i < 100; i++ {
		fmt.Printf("You say: %q\n", <-msg)
	}
}
