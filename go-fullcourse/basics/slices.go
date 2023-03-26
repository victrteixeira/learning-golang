package basics

import "fmt"

func Slice_Declaration() (string) {
	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	res1 := fmt.Sprintf("Length: %v\n", len(a))
	res2 := fmt.Sprintf("Capacity: %v", cap(a))

	return fmt.Sprint(res1) + fmt.Sprintln(res2)
}

func Slice_SliceByRange() (string) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th, and 6th elements
	a[5] = 42
	// the first number is inclusive, and the second number is exclusive

	return fmt.Sprintf("A: %v\n", a) + fmt.Sprintf("B: %v\n", b) + fmt.Sprintf("C: %v\n", c) + fmt.Sprintf("D: %v\n", d) + fmt.Sprintf("E: %v\n", e)
}

func Slice_BuiltInMakeFunction() (string){
	a := make([]int, 3, 100)
	sliceValues := fmt.Sprint(a)
	lengthValue := fmt.Sprintf("Length: %v", len(a))
	capabilityValue := fmt.Sprintf("Capability: %v", cap(a))

	return fmt.Sprintln(sliceValues) + fmt.Sprintln(lengthValue) + fmt.Sprintln(capabilityValue)
}

func Slice_AppendToSlice() (string) {
	a := []int{}
	sliceValues := fmt.Sprint(a)
	lengthValue := fmt.Sprintf("Length: %v", len(a))
	capabilityValue := fmt.Sprintf("Capability: %v", cap(a))

	a = append(a, 1)
	sliceValues2 := fmt.Sprint(a)
	lengthValue2 := fmt.Sprintf("Length: %v", len(a))
	capabilityValue2 := fmt.Sprintf("Capability: %v", cap(a))

	a = append(a, 2, 3, 4, 5)
	sliceValues3 := fmt.Sprint(a)
	lengthValue3 := fmt.Sprintf("Length: %v", len(a))
	capabilityValue3 := fmt.Sprintf("Capability: %v", cap(a))
	
	return fmt.Sprintln(sliceValues) + fmt.Sprintln(lengthValue) + fmt.Sprintln(capabilityValue) + fmt.Sprintln(sliceValues2) + fmt.Sprintln(lengthValue2) + fmt.Sprintln(capabilityValue2) + fmt.Sprintln(sliceValues3) + fmt.Sprintln(lengthValue3) + fmt.Sprintln(capabilityValue3)
}

func Slice_AppendSliceToSlice() (string) {
	a := []int{}
	sliceValue := a
	fmt.Println(a)
	lengthValue := fmt.Sprintf("Length: %v", len(a))
	capabilityValue := fmt.Sprintf("Capability: %v", cap(a))

	fmt.Println("Appending...")
	a = append(a, []int{1,2,3,4,5}...) // spread operator (in javascript at least) -> ...

	sliceValue2 := a
	fmt.Println(sliceValue2)
	lengthValue2 := fmt.Sprintf("Length: %v", len(a))
	capabilityValue2 := fmt.Sprintf("Capability: %v", cap(a))

	return fmt.Sprintln(sliceValue) + fmt.Sprintln(lengthValue) + fmt.Sprintln(capabilityValue) + fmt.Sprintln(sliceValue2) + fmt.Sprintln(lengthValue2) + fmt.Sprintln(capabilityValue2)
}

func Slice_PoppingElementsOffSlice() (string) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original Slice: %v\n", a)

	b := a[1:len(a)-1] // that remove the last number of a slice -> number 5 of this slice
	fmt.Printf("Removing the end and beginning of slice: %v\n", b)

	c := append(a[:2], a[3:]...) // removes the middle of a slice -> number 3 in that case
	return fmt.Sprintf("Removing the middle of slice: %v\n", c)
}