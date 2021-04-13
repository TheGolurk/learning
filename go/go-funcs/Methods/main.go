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

	if err := ReadFullFile(); err == io.EOF {
		println("successfully read file")
	} else if err != nil {
		fmt.Printf("Something bad ocurred %s \n", err)
	}
}

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("a panic ocurred but it is ok")
		} else if p != nil {
			panic("an unexpected error ocurred and we do not want to recover")
		}
	}()

	for {
		value, readerErr := r.Read([]byte("text that does nothing"))
		if readerErr == io.EOF {
			break
		} else if readerErr != nil {
			err = readerErr
			return
		}

		println(value)

	}
	return nil
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

type SimpleReader struct {
	count int
}

var errCatastrophicReader = errors.New("something catastrophic occurre in the reader")

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	if br.count == 2 {
		panic(errCatastrophicReader)
	}
	if br.count > 3 {
		return 0, io.EOF //errors.New("bad robot") //io.EOF
	}
	br.count += 1
	return br.count, nil
}

func (br *SimpleReader) Close() error {
	println("closing reader")
	return nil
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
		panic("an invalid math expression was used")
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}
