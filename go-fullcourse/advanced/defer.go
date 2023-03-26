package advanced

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Defer_Concept(){
	//fmt.Println("start")
	// defer fmt.Println("middle") 
	//fmt.Println("end")

	// The defer functions executes after the end of function, but before his return

	fmt.Println()

	defer fmt.Println("start 2 ") // Third
	defer fmt.Println("middle 2") // Second
	defer fmt.Println("end 2") // First

	// When there are more than one defer order, the function executes a LIFO order or last in first out (LIFO)
}

func Defer_RealExample(){
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // Without defer function here, this application will get an error, because we'll try to read after the request body has been closed.
	
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}