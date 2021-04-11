package main

import (
	"github.com/TheGolurk/go-functions/simplemath"
)

func main() {
	sv := simplemath.NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	println(sv.String())

	func() {
		println("hello")
	}()
}
