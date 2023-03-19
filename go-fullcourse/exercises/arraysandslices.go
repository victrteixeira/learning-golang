package exercises

import "fmt"

func Array_Example1() string {
	grades := [3]int{32, 21, 42} // declaration
	return fmt.Sprintf("Grades: %v\n", grades)
}
