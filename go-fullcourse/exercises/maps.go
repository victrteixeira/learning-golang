package exercises

import "fmt"

func Map_Declaration() (string) {
	// declaration
	statePopulation := make(map[string]int) // or
	statePop := make(map[string]int, 10) // or
	statePopulations := map[string]int {
		"California": 39250017,
		"Ohio": 11614373,
	}

	return fmt.Sprintln(statePopulations, statePopulation, statePop)

	// Map is a strong typed key-value pair.
}

func Map_Manipulation() (string){
	statePopulations := map[string]int {
		"California": 39250017,
		"Ohio": 11614373,
	}

	statePopulations["Georgia"] = 10310371 // It's possible to add a new key-value pair to the map with this syntax.

	delete(statePopulations, "California") // It's possible to delete a key-value pair from a map with built-in function "delete" -> the first is the map, the second is the value, in that case "California", but it could be "Ohio" or "Georgia".

	georgia := statePopulations["Georgia"] // It's possible to get any value from a map, just by passing his key.
	fmt.Printf("Georgia Value: %v\n", georgia)

	pop, ok := statePopulations["Oho"] // When a value that doesn't exist inside a map is called, it return 0, but to properly see whether the value is ou is not inside a map.

	return fmt.Sprintf("Pop: %v,\nOk: %v,\n", pop, ok)
}