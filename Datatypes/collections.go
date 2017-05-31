package main

import "fmt"

func main() {
	Array1()
	Array2()
	slice()
	make_demo()
	map_demo()
}
func blank() {
	println("                                            ")
}

func Array1() {
	blank()
	println("Basic array information")

	//Array example one 
	myArray := [5]int{}
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 20
	myArray[3] = 40
	myArray[4] = 50

	fmt.Println(myArray)
}

func Array2() {
	blank()
	println("Short hand array")

	myArray := [...]int{60, 70, 80, 90, 100}
	fmt.Println(myArray)
}

func slice() {
	println("Printing slice information")
	myArray := [...]int{60, 70, 80, 90, 100}

	myslice := myArray[:]

	myslice = append(myslice, 20)
	fmt.Println(myArray)
	fmt.Println(myslice)

	fmt.Println(len(myArray))
	fmt.Println(len(myslice))
}

func make_demo() {
	blank()
	println("Make function with array")

	myslice := make([]int, 100)
	myslice[10] = 10
	myslice[21] = 21
	myslice[33] = 33
	myslice[98] = 99
	myslice[99] = 100
	fmt.Println(myslice)
}

func map_demo() {
	blank()
	println("Map function demo with make")

	myMap := make(map[int]string)
	fmt.Println(myMap);
	myMap[20] = "Twenty"
	myMap[1] = "One"
	fmt.Println(myMap);
}
