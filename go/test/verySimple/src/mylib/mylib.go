package mylib

import "fmt"

func messageWriter(greeing, name string) string {
	return fmt.Sprintf("%v, %v", greeing, name)
}
