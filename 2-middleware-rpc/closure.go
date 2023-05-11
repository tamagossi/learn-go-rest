package middleware_rpc

import "fmt"

func generator() func() int {
	var i = 0

	return func() int {
		i++
		return i
	}
}

func Closure() {
	numGenerator := generator()

	for i := 0; i < 5; i++ {
		fmt.Print(numGenerator(), "\t")
	}
}
