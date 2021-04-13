package main

import (
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 4; i++ {
		id := rnd.Intn(4) + 1
		println(id)
	}
}
