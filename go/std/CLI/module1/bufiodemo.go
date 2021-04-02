package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "/quit" {
			fmt.Println("Exiting")
			os.Exit(3)
		} else {
			fmt.Println("You typed" + scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
