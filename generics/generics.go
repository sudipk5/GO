package main

import (
	"fmt"
)
func printslice[T any](items []T) {

	for _, item := range items {

		fmt.Println(item)
	}
}
func advensprintslice[T int | string](items []T) {

	for _, item := range items {

		fmt.Println(item)
	}
}



type sta[T int | string]struct{

	element[]T
}





func main() {

	// num := []int{1, 2, 3}
	// advensprintslice(num)

	mytype:=sta[int]{
		element:[]int{1,2,3,4},
	}
   fmt.Println(mytype)
}
