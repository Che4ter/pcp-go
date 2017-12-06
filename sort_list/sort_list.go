package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	unsortedList := readList()
	sortedList := sortList(unsortedList)
	fmt.Println(sortedList)
}

func readList() []int {
	fmt.Println("Enter numbes separated with a space:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	
	/* Remove LF and CR */
	input = strings.Replace(input, "\r", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	
	unformatedList := strings.Split(input, " ")
	
	unsortedList := make([]int, len(unformatedList))
	for i, n := range unformatedList {
		unsortedList[i], _ = strconv.Atoi(n)
	}
	
	return unsortedList
}

func sortList(unsortedList []int) []int {	
	for i, n := range unsortedList {
		if i < len(unsortedList) - 1 {
			if n > unsortedList[i+1] {
				unsortedList[i] = unsortedList[i+1]
				unsortedList[i+1] = n
			}
		}
	}
	
	return unsortedList
}