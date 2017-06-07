package main

import (
	"fmt"
)

func main() {

	//Basic Slice creation
	var s []int
	s = make([]int, 3, 10)
	s[0] = 1
	s[1] = 10
	s[2] = 100

	//Short hand syntax
	s2 := []int{1, 10, 100, 1000}

	//Slicing the slice
	slice := s2[1:]

	//appending the slice
	slice = append(slice, 10000)


	//Deleting a slice 
	nSlice := append(s2[1:] , slice[:2]...)


	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(slice)
	fmt.Println(nSlice)
	fmt.Println("Demo for slices")
}

