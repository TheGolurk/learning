package main

import (
	"flag"
	"fmt"
)

func main() {
	archPtr := flag.String("arch", "x86", "CPU Type")

	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("Running in 32 bits")
	}

	fmt.Println("Process complete")
}
