package main

import (
	"fmt"
	"time"
)

func talking_basic(receiver string) {
	for i := 0; ; i++ {
		fmt.Println("Hello ", receiver, ":", i)
		<-time.After(time.Millisecond * 200)
	}
}

func main() {
	fmt.Println("Start talking")
	go talking_basic("Alan")
	go talking_basic("Philipp")

	<-time.After(time.Second * 5)
}
