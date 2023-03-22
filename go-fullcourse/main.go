package main

import (
	"fmt"
	"full-course/exercises"
)

func main() {
	//res := exercises.Struct_Declaration()
	//fmt.Print(res)
	
	//resNumber, resString, resSlice := exercises.Struct_Manipulating()
	//fmt.Printf("Number: %v,\nString: %v,\nSlice: %v\n", resNumber, resString, resSlice)

	//exercises.Struct_Composition()

	//exercises.Struct_Tags()
	fmt.Printf("This is the entire struct: %v\n", exercises.Struct_FunctionWithAnonymousStruct())
	fmt.Printf("This is only the company field: %v\n", exercises.Struct_FunctionWithAnonymousStruct().Company)
	fmt.Printf("This is only the phone field: %v\n", exercises.Struct_FunctionWithAnonymousStruct().Phone)
}
