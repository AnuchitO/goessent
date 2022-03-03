package main

import (
	"fmt"

	"github.com/AnuchitO/goessent/say"
)

func main() {
	s := "Hello Gopher"

	fmt.Println(s)

	a, b := 1, "Hellllll"
	fmt.Println(a, b)

	h, a := say.Say(b)
	fmt.Println(h, a)
}
