package end

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