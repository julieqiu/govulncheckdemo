// This program takes language tags as command-line
// arguments and parses them.

package main

import (
	"fmt"
	"os"

	"golang.org/x/text/language"
)

func main() {
	// Calls foo which eventually calls language.Parse.
	A()

	// Also calls foo which eventually calls language.Parse.
	B()

	// Calls language.Parse directly.
	C()

	// Calls foobar which eventually calls language.MustParse (different
	// symbol, same report)
	D()
}

func A() {
	foo(os.Args[1:])
}

func B() {
	foo(os.Args[1:])
}

func C() {
	_, _ = language.Parse("")
}

func D() {
	foobar()
}

func foo(args []string) {
	for _, arg := range args {
		tag, err := language.Parse(arg)
		if err != nil {
			fmt.Printf("%s: error: %v\n", arg, err)
		} else if tag == language.Und {
			fmt.Printf("%s: undefined\n", arg)
		} else {
			fmt.Printf("%s: tag %s\n", arg, tag)
		}
	}
}

func foobar() {
	language.MustParse("")
}
