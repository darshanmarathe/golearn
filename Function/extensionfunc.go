package main

import (
	"fmt"
)




type Person struct {
	FirstName string
	MiddleName string
	LastName string
}


func (person *Person) FullName(title string){
	fmt.Println(title + " " +
				person.FirstName + " "  + 
				person.MiddleName + " " + 
				person.LastName)
}

func main(){
	var me = Person{"Darshan" , "Narayan" , "Marathe"}
	me.FullName("Mr")
}