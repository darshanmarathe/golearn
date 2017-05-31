package main

import (
	"fmt"
)

type name struct{
	firstName string 
	middleName string 
	lastName string
} 

func main(){

	var fullname = 	name{"Darshan" , "Narayan" , "Marathe"}

	var wifename = name{middleName:"Darshan" , lastName:"Marathe" , firstName:"aditi"}

	fmt.Println(fullname)
	fmt.Println(wifename)
}