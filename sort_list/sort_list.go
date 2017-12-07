package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	unsortedList := readList()
	asc := readAsc()
	sortedList := sortList(unsortedList, asc)
	fmt.Println(sortedList)
}

func readAsc() bool {
	fmt.Println("Type ASC for ascending:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	
	return strings.ContainsAny("ASC", strings.ToUpper(input))
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

func sortList(unsortedList []int, asc bool) []int {
	if asc == true {
		sort.Sort(sort.IntSlice(unsortedList))
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(unsortedList)))
	}
	
	return unsortedList
}