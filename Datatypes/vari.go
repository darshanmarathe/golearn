package main

import "fmt"

func main() {
	var noOfyears int
	noOfyears = 30

	var gotPercent float32 = 47.60

	myName := "Darshan Marathe"

	mycomplex := complex(2, 3)

	fmt.Println(noOfyears)
	fmt.Println(gotPercent)
	fmt.Println(myName)

	fmt.Println(mycomplex)
	fmt.Println(real(mycomplex))
	fmt.Println(imag(mycomplex))
	MoreSamples()
}

func MoreSamples() {
	var message string = "hello more examles"
	var a, b, c int = 1, 2, 3


	fmt.Println(message, a, b, c)

	//type infer
	var x, y, z = "string"  , false , 20

	fmt.Println(x , y, z)

}
