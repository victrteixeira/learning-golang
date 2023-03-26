package advanced

import "fmt"

// This section only have information about things on function that I don't already know.

func Functions_Pointers(){
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println("############ The function ended ############")
	fmt.Println(name)
	// To see the differences either with and without pointers, just remove his syntax -> "&" and "*" from the fields.
}

func sayGreeting(greeting, name *string){
	fmt.Println(*greeting, *name)
	fmt.Println("############ Before the assign value ############")

	*name = "Ted"
	fmt.Println(*name)
}

func Functions_VariadicParameters(){
	sum(1,2,3,4,5,6,7,8,9)
	sum_secondExample("The sum on the second example is:", 1,2,3,4,5,6,7,8,9)
}

func sum(values ...int){ // The variadic parameter told the go runtime to take in all of the last arguments that are passed in and wrap them up into a slice.
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("the sum is:", result)
}

func sum_secondExample(msg string, values ...int){ // It's possible to use variadic parameters with others, but a function could only have one and it must be at the end.
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

func Functions_ReturnAsPointer(){
	s := returnAsPointer(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s)
}

func returnAsPointer(values ...int) (*int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	return &result
}

func Functions_NamedReturnValues(){
	s := namedReturnValues(1, 2, 3, 4, 6)
	fmt.Println("The sum is", s)
}

func namedReturnValues(values ...int) (result int){
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func Functions_MultipleReturnsAndError(){
	d, err := divide(5.0, 9.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}

func divide(a, b float64) (float64, error){
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func Functions_AnonymousFunctions(){
	func() { 
		fmt.Println("Hello Go!") 
	}()

	// An another approach to this

	f := func() {
		fmt.Println("Hello Go 2!")
	}
	
	f()
	
	// The same function above, without the syntax sugar of go compiler
	
	var f2 func() = func(){
		fmt.Println("Hello Go 3!")
	}
	f2() // It works because Go compiler understand functions as types.
}

func Functions_AnonFuncRealExample(){
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error){
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		} else {
			return a / b, nil
		}
	}

	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}

func Functions_Methods(){
	g := greeter {
		greeting: "Hello",
		name: "Go",
	}
	
	g.greet()
}

type greeter struct {
	greeting string
	name string
}

func (g greeter) greet(){
	fmt.Println(g.greeting, g.name)
}