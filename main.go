package main

import (
	"fmt"

	"github.com/stephenkingston/go-getting-started/anotherlib"
	"github.com/stephenkingston/go-getting-started/anotherlib/blip"
	"github.com/stephenkingston/go-getting-started/mylib"
	"github.com/stephenkingston/go-getting-started/mylib/advanced"
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
