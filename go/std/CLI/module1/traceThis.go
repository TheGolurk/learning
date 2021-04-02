// Go tool: go tool trace trace.out
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal("we did not create a file", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("failed to close file", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatal("failed to start a trace", err)
	}

	defer trace.Stop()

	AddRandomNumbers()
}

func AddRandomNumbers() {
	fnum := rand.Intn(100)
	snum := rand.Intn(100)

	time.Sleep(2 * time.Second)

	result := fnum * snum

	fmt.Println(result)
}
