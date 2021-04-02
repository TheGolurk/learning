package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("name")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	fmt.Println("Running in" + runtime.Version() + runtime.GOOS)

}
