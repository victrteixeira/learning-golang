package advanced

import "fmt"

func Interfaces_Concept(){
	var w writer = consoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type writer interface {
	Write([]byte) (int, error)
}

type consoleWriter struct {}

func (cw consoleWriter) Write(data []byte) (int, error){
	n, err := fmt.Println(string(data))
	return n, err
}

func Interfaces_WithPrimitiveTypes(){
	myInt := intCounter(0)
	var inc incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type incrementer interface {
	Increment() (int)
}

type intCounter int

func (ic *intCounter) Increment() (int){
	*ic++
	return int(*ic)
}