package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) string {
	rawExpression := strings.Fields(text)
	cleanExpression := strings.Join(rawExpression, "")
	return cleanExpression
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
}
