package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	X int
	Y int
}

func main() {
	for _, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(v)
		if v%2 == 0 {
			fmt.Println("Hello World!")
		}

		foo := Foo{}
		foo.X = 123
		foo.Y = 456

		bar := Bar(foo)
		byte, err := json.Marshal(bar)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(byte))

	}
}

func Bar(foo Foo) Foo {
	foo.X = foo.X * 100
	foo.Y = foo.Y * 100

	return foo
}
