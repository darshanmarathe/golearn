package maths

import "fmt"

func Add(in ...int) (output int) {
	for _, num := range in {

		output += num
	}

	fmt.Println(output)
	return
}

func Multiply(in ...int)(output int){
	for _, num := range in {

	output *= num
	}

	fmt.Println(output)
	return
}
