package advanced

import "fmt"

func Pointers_Concept(){
	var a int = 42
	var b *int  = &a // This variable 'b' is pointing to 'a' variable. That's the pointer's syntax in Go.

	fmt.Println(a, b) // Here, it's printed the value and the memory location of 'a'
	fmt.Println(&a, b) // It can be visible here, showing that either variables have the same memory location.

	fmt.Println(a, *b) // The star before variable is called 'dereferences'. It's means that when there's it before a variable, the underlying value will be exposed.
	
	a = 27
	fmt.Println(a, *b) // Even when the 'primary' variable has changed his value, the pointer will always follow the changes.
	
	*b = 14
	fmt.Println(a, *b) // The opposite is true as well. When a dereferenced pointer change his value, the 'primary' variable will change too.
}

func Pointers_WithStructs(){
	var ms *myStruct
	//ms = &myStruct{foo: 42} // It's possible to do the same thing with the built-in "new" function.
	ms = new(myStruct) // On this approach it isn't possible to have initialization syntax as above.

	(*ms).foo = 42 // This is the standard syntax to initalize this struct, it'll be necessary to use the parentheses to derefence the pointers.
	fmt.Println((*ms).foo)

	// However, Go compiler has a syntax sugar to make this more simple. The example below runs the same way as the above, but now with less verbose syntax.

	var mp *myStruct
	mp = new(myStruct)
	mp.foo = 44

	fmt.Println(mp.foo)
}

func Pointers_VariablesAssignOneToAnother(){
	a := []int { 1, 2 , 3}
	b := a
	fmt.Println(a, b)
	
	a[1] = 42
	fmt.Println(a, b)

	// Slices and maps have underlying pointers inside them. Because of it, when a copy is done into another variable, and that copy is a copy of a map or slice, the values will be changed in all copies.
}

type myStruct struct {
	foo int
}