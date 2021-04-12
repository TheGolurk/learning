package main

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/TheGolurk/go-functions/simplemath"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	sv := simplemath.NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	println(sv.String())

	func() {
		println("hello")
	}()

	addExpr := mathExpression(AddExpr)
	println(addExpr(2, 3))

	fmt.Printf("%f\n", double(3, 4, mathExpression(AddExpr)))

	p2 := powerOfTwo()
	value := p2()
	println(value)

	ReadSomething()
}

func ReadSomething() error {
	var r io.Reader = BadReader{errors.New("mynosense reader")}
	if _, err := r.Read([]byte("test")); err != nil {
		fmt.Printf("an error ocurred %s \n", err)
		return err
	}
	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		return func(f1, f2 float64) float64 {
			return 0
		}
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}
