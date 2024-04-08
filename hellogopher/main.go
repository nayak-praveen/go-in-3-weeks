package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
	"github.com/nayak-praveen/stringutils"
)

const location = "Remote"

var name string

func main() {

	name = "Praveen"
	from := `Phoenix,Arizona`
	var n int

	var proverb = "Undefined"
	if p, err := proverbs.Nth(4); err == nil {
		proverb = p.Saying
	}

	fmt.Printf("Hello, fellow %s Gophers!\n", location)
	fmt.Printf("My name is %s and I'm from %s.\n", stringutils.ToUpperCase(name), stringutils.ToLowerCase(from))
	fmt.Printf("By the time %d o'clock EST comes around, we'll know how to code in Go!\n", n)
	fmt.Printf("Here's a Go proverb: %s\n", proverb)
	fmt.Println("Let's get started!")
}
