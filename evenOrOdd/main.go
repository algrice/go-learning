package main

import "fmt"

func main() {
	// low, high := 0, 10
	// s := make([]int, high-low+1)
	// for i := range s {
	// 	s[i] = i
	// }
	// fmt.Println(s)

	// for _, n := range s {
	// 	if n%2 == 0 {
	// 		fmt.Println("Even")
	// 	} else {
	// 		fmt.Println("Odd")
	// 	}
	// }

	low, high := 0, 10
	s := make([]int, high-low+1)
	for i := range s {
		s[i] = i
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

}
