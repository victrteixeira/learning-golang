package exercises

import "fmt"

func oneReturn() string {
	x := 42
	y := "James Bond"
	z := true

	res := fmt.Sprintf("%v, %v, %v\n", x, y, z)

	return res
}

func MultipleReturns() {
	x := 24
	y := "James Bond 2"
	z := true

	fmt.Print(x, "\n")
	fmt.Print(y, "\n")
	fmt.Print(z, "\n")
}

func AccessingOneReturn() string {
	return oneReturn()
}
