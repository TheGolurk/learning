package main

import (
	"fmt"

	"github.com/TheGolurk/go-functions/simplemath"
)

func main() {
	fmt.Println("vim-go")
	sv := simplemath.NewSemanticVersion(1, 2, 3)
	println(sv.String())
}
