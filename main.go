package main

import (
	"fmt"
)

type Color string

const (
	Red    Color = "red"
	Blue   Color = "blue"
	Yellow Color = "yellow"
)

type Foo string

const (
	FooRed Foo = "f00-red"
)

func main() {
	fmt.Println(Color(Red))
	fmt.Println(Foo(FooRed))
}
