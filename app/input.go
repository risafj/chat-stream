package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputFromCommandLine() string {
	var input string
	fmt.Printf("Message: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}
