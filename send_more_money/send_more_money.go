package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	equation := readEquation()
	equation = "SEND + MORE = MONEY"
	summandA := getSummand(equation, true)
	summandB := getSummand(equation, false)
	sum := getSum(equation)

	fmt.Println(summandA, "+", summandB, "=", sum)
	solve(summandA, summandB, sum)
}

func readEquation() string {
	fmt.Println("[SEND + MORE = MONEY]")
	fmt.Print("Enter equation: ")
	reader := bufio.NewReader(os.Stdin)
	equation, _ := reader.ReadString('\n')

	return strings.ToUpper(equation)
}

func getSummand(equation string, first bool) string {
	var summand string

	if first == true {
		summand = strings.Split(equation, "+")[0]
	} else {
		summand = strings.Split(equation, "+")[1]
		summand = strings.Split(summand, "=")[0]
	}

	return strings.Replace(summand, " ", "", -1)
}

func getSum(equation string) string {
	sum := strings.Split(equation, "=")[1]

	return strings.Replace(sum, " ", "", -1)
}

func solve(summandA, summandB, sum string) {
	weightedA := weighExpression(summandA)
	weightedB := weighExpression(summandB)
	weightedSum := weighExpression(sum)

	/* exit if sum length is shorter than a */
	if len(weightedA) > len(weightedSum) {
		fmt.Println("Sum needs to be longer than Summand")
		os.Exit(3)
	}

	zero := []byte{65}
	for len(weightedSum) > len(weightedA) {
		weightedA = append(zero, weightedA...)
	}
	for len(weightedSum) > len(weightedB) {
		weightedB = append(zero, weightedB...)
	}

	/* map letters to numbers */
	for _, n := range weightedA {
		allDistinct(n)
	}
	for _, n := range weightedB {
		allDistinct(n)
	}
	for _, n := range weightedSum {
		allDistinct(n)
	}

	/* reassign summand numbers */
	for i, n := range weightedSum {
		indexA := weightedA[i]
		indexB := weightedB[i]
		weightedA[i] = byte(distinct[indexA])
		weightedB[i] = byte(distinct[indexB])
		weightedSum[i] = byte(distinct[n])
	}

	i := len(weightedA) - 1
	for i > 0 {
		valSum := weightedA[i] + weightedB[i]
		fmt.Print(distinct[valSum])
		i--
	}
	fmt.Println()

	/* Print neatly aligned */
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, summandA, "\t", weightedA)
	fmt.Fprintln(w, summandB, "\t", weightedB)
	fmt.Fprintln(w, sum, "\t", weightedSum)
	w.Flush()
}

var distinct = make(map[byte]int)
var index int

func allDistinct(n byte) {
	if _, ok := distinct[n]; ok {
		/* do nothing */
	} else {
		distinct[n] = index
		index++
	}
}

func weighExpression(expression string) []byte {
	weightedExpression := []byte(expression)
	for i, c := range weightedExpression {
		if c == 13 {
			weightedExpression = weightedExpression[:i]
		}
	}

	return weightedExpression
}
