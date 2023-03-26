package basics

import "fmt"

func Array_SimpleDeclaration() string {
	grades := [3]int{32, 21, 42} // declaration
	return fmt.Sprintf("Grades: %v\n", grades)
}

func Array_ManagingSize() string {
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Victor"
	students[1] = "Lisa"
	students[2] = "Arnold"
	fmt.Printf("Student #1: %v\n", students[0])
	return fmt.Sprintf("Number of Students: %v\n", len(students))
}

func Array_MatrixDeclaration() string {
	var matrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Printf("Matrix here we go:\n%v\n", matrix)

	var matrixAlternative [3][3]int
	matrixAlternative[0] = [3]int{1, 0, 0}
	matrixAlternative[1] = [3]int{0, 1, 0}
	matrixAlternative[2] = [3]int{0, 0, 1}
	return fmt.Sprintln(matrixAlternative)
}

func Array_AsValues() string {
	a := [...]int{1, 2, 3}
	b := &a // Without pointer, Go will copy the entire array over, what could let the program slow.
	b[1] = 5
	res := fmt.Sprintln(a) + fmt.Sprintln(b)
	return res
}
