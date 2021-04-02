package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	t := "the go standard library"
	tim := timed(properTitle).(func(string) string)
	ntt := tim(t)
	fmt.Println(ntt)
	timedtoo := timed(num).(func(int) int)
	fmt.Println(timedtoo(1))
}

func timed(f interface{}) interface{} {
	rf := reflect.TypeOf(f)

	if rf.Kind() != reflect.Func {
		panic("expected func")
	}
	vf := reflect.ValueOf(f)
	wrapper := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Println(runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapper.Interface()
}

func properTitle(in string) string {
	words := strings.Fields(in)
	smallWords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallWords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func num(a int) int {
	time.Sleep(1 * time.Second)
	return a * 2
}
