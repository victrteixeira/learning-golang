package end

import (
	"fmt"
	"net/http"
)

func Panic_Concept(){
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
	// Panic is nothing more than an exception. It's the same concept.
}

func Panic_SimpleSyntaxExample(){
	fmt.Println("Start")
	panic("something bad happened")
	fmt.Println("End")
}

func Panic_RealExample(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Hello Go!"))
	})
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func Panic_WithDefer(){
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
	// Although defer currently runs after the function ends, panic will run only after defer statement.

	// The order is: 
	//main function -> handle defer func -> handle any panic statements -> and then handle return values
}