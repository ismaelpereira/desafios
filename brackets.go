package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	test := "[([])()({})]"
	var brackets []byte
	var result bool
	var err error
	for i := 0; i < len(test); i++ {
		if test[i] == 40 || test[i] == 91 || test[i] == 123 {
			brackets = append(brackets, test[i])
			continue
		}
		if test[i] == 41 &&
			len(brackets) != 0 &&
			brackets[len(brackets)-1] == 40 {
			brackets = removeFromArray(brackets, len(brackets)-1)
			continue
		} else if test[i] == 93 &&
			len(brackets) != 0 &&
			brackets[len(brackets)-1] == 91 {
			brackets = removeFromArray(brackets, len(brackets)-1)
			continue
		} else if test[i] == 125 &&
			len(brackets) != 0 &&
			brackets[len(brackets)-1] == 123 {
			brackets = removeFromArray(brackets, len(brackets)-1)
			continue
		} else {
			fmt.Println(result)
			return err
		}
	}
	if len(brackets) == 0 {
		result = true
	}
	fmt.Println(result)
	return err
}

func removeFromArray(a []byte, index int) []byte {
	return append(a[:index], a[index+1:]...)
}

/*
ASCII TABLE
( = 40
) = 41
[ = 91
] = 93
{ = 123
} = 125
*/

/*
[[[((({{{}}})))]]] TRUE
[()]{}[()()]() TRUE
[]()()((([]))) TRUE
[]()()(((([]))) FALSE
[(]){}{[()()]()} FALSE
[()]{}{[()()]()] FALSE
( FALSE
[ FALSE
{ FALSE
) FALSE
] FALSE
} FALSE
[(]) FALSE
[]()()(((([]))) FALSE
[](){{{[]}}} TRUE
[]{}()[][][] TRUE
([()][][{}]) TRUE
[([])()({})] TRUE
*/
