package main

import "fmt"

func main(){
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
}