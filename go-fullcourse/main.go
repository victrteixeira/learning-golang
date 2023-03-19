package main

import (
	"fmt"
	"full-course/exercises"
)

func main() {
	fmt.Print("Accessing OneReturn private function\n\n")
	res := exercises.AccessingOneReturn()
	fmt.Printf("%v\n", res)

	fmt.Print("Done\n\n")
}
