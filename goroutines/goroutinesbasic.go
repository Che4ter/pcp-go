package main

import (
	"fmt"
	"time"
)

func talking(receiver string) {
	for i := 0; ; i++ {
		fmt.Println("Hello ", receiver, ":", i)
		<-time.After(time.Millisecond * 200)
	}
}

func main() {
	fmt.Println("Start talking")
	go talking("Alan")
	go talking("Philipp")

	<-time.After(time.Second * 5)
}
