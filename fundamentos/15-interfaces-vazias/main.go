package main

import "fmt"

func main() {
	var x interface{} = 1000
	var y interface{} = "Hello World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}
