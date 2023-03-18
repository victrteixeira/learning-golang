package main

import (
	"fmt"
	printandvar "full-course/exercises"
)

func main() {
	fmt.Print("Accessing OneReturn private function\n\n")
	res := printandvar.AccessingOneReturn()
	fmt.Printf("%v\n", res)

	fmt.Print("Done\n\n")

	printandvar.MultipleReturns()
}
