package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func readPerson() string {
	fmt.Print("Check person: ")
	reader := bufio.NewReader(os.Stdin)
	person, _ := reader.ReadString('\n')
	
	/* remove LF and CR (linefeed and carriage return) */
	return strings.ToLower(person[:len(person) - 2])
}

var brother = make(map[string]string)
func main() {
	initBrother()
	person := readPerson()
	findBrother(person)
}

func initBrother() {
	brother["alan"] = "philipp"
	brother["philipp"] = "alan"
}

func findBrother(person string) {
	fmt.Println("The brother of", person, "is", brother[person])
}