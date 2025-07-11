package main

import (
	"fmt"
	"go-getting-started/anotherlib"
	"go-getting-started/anotherlib/blip"
	"go-getting-started/mylib"
	"go-getting-started/mylib/advanced"
)

func getRandom() int {
	return 4
}

func main() {
	fmt.Println("Hello World")

	fmt.Println(anotherlib.Hello("Stephen"))
	fmt.Println(anotherlib.Bye("Stephen"))

	fmt.Println(mylib.Add(1, 2))
	fmt.Println(mylib.Sub(1, 2))
	fmt.Println(mylib.Mul(1, 2))

	fmt.Println(blip.Blip())

	fmt.Println(getRandom())

	advanced.AdvancedStuff()
}
