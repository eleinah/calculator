package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) string {
	fields := strings.Fields(text)
	final := strings.Join(fields, "")
	return final
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter expression >> ")
		reader.Scan()

		//exp := reader.Text()
	}
}
