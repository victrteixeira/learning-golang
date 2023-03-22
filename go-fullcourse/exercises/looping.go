package exercises

import (
	"fmt"
)

func Loop_SimpleDeclaration() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("First for passed, now the second one.")

	for i := 0; i < 10; i = i+2 {
		fmt.Println(i)
	}

	fmt.Println("Second for passed, now the third one.")

	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}

func Loop_CreatingLogic(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2 * i + 1
		}
	}
}

func Loop_ShortDeclaration(){
	i := 0
	// for ; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for ; i < 5; {
	// 	fmt.Println(i)
	// 	i++
	// }

	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func Loop_InfiniteLoops(){
	i := 0
	for {
		fmt.Println(i)
		i++
		if i > 100 { break }
	}
}

func Loop_NestedLoopAndHowToBreakIt(){
	Loop: 
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i * j >= 3 { break Loop}
		}
	}
}

func Loop_Collections(){
	s := [3]int{1, 2, 3}

	for k, v := range s {
		fmt.Printf("Key: %v\nValue: %v\n", k, v)
	}

	fmt.Println()

	b := "hello Go!"
	for _, v := range b {
		fmt.Println(string(v))
	}
}