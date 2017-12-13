package main

import (
	"fmt"
	"log"
	"os/exec"
	"main"
)

func main() {
	cmd := GetCommand()
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}