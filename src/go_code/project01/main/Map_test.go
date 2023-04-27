package main

import (
	"fmt"
	"testing"
)

func Test_function1(t *testing.T) {
	var tests = []struct {
		name      string
		funcnames funcname
	}{
		{name: "function", funcnames: funcname{"func"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{
				tt.funcnames.function1()
				tt.funcnames.function2()
				name := tt.funcnames.name
				fmt.Println(name)
			}
		})
	}
}
