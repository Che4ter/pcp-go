package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readPerson() string {
	fmt.Print("Check person: ")
	reader := bufio.NewReader(os.Stdin)
	person, _ := reader.ReadString('\n')
	
	person = strings.Replace(person, "\r", "", -1)
	person = strings.Replace(person, "\n", "", -1)

	/* remove LF and CR (linefeed and carriage return) */
	return strings.ToLower(person)
}

var brother = make(map[string]string)

func main() {
	initBrother()
	person := "alan"
	findBrother(person)
}

func initBrother() {
	brother["alan"] = "philipp" // HL
	brother["philipp"] = "alan"
}

func findBrother(person string) {
	fmt.Println("The brother of", person, "is", brother[person])
}
