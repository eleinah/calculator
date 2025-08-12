package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func cleanInput(input string) string {
	rawExpression := strings.Fields(input)
	cleanExpression := strings.Join(rawExpression, "")
	return cleanExpression
}

func parseExpression(raw string) []string {
	return []string{}
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter expression >> ")
		reader.Scan()

		//exp := reader.Text()
	}
}

func main() {
	fmt.Println("Hello, world!")
	mth := math.Pow((1+2*3), 4)*12 - 1
	fmt.Printf("answer = %.2f\n", mth)
}
