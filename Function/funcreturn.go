package main

import (
	"fmt"
)

func main() {
	num := Add(10, 30, 20, 40)
	count, num1 := AddandCount(10, 30, 20, 40)
	count1, num2 := AddandCount_Namedreturn(10, 30, 20, 40)
	_, num3 := AddandCount_Namedreturn(10, 30, 20, 40)

	println(num)
	fmt.Println(count, num1)
	fmt.Println(count1, num2)
	fmt.Println(num3)
	fmt.Println(Concat("Hello" , "I am" , "Darshan" , "Marathe" ))
}

func Concat(s1,s2,s3,s4 string)string{
	space := " "
	return s1 + space + s2 + space + s3 + space + s4
}

func Add(num ...int) int {
	result := 0
	for _, v := range num {
		result += v
	}
	return result
}

func AddandCount(nums ...int) (int, int) {
	result := 0
	for _, v := range nums {
		result += v
	}
	return len(nums), result
}


//named return values
func AddandCount_Namedreturn(nums ...int) (count int, result int) {

	for _, v := range nums {
		result += v
	}
	count = len(nums)
	return
}
