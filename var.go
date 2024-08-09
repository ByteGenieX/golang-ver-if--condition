package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	name = "Neo"
	var x = 22
	y := 36.45

	fmt.Println("-------------------")
	fmt.Println(name)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
	fmt.Println("-------------------")
}
