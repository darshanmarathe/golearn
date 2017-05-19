package main

import (
	"fmt"
)


func main(){
   	num := 	Add(10 , 30 , 20 , 40)
	count , num1 := 	AddandCount(10 , 30 , 20 , 40)
  count1 , num2 := 	AddandCountV2(10 , 30 , 20 , 40)
  
   	println(num)
	fmt.Println(count, num1)
	fmt.Println(count1, num2)
}


func Add(num ...int) int {
	result :=0
	for _, v := range num {
		result += v
	}
	return result
}


func AddandCount(nums ...int) (int , int) {
	result :=0
	for _, v := range nums {
		result += v
	}
	return len(nums) , result
}

func AddandCountV2(nums ...int) (count int , result int) {
	
	for _, v := range nums {
		result += v
	}
	count = len(nums)
	return
}