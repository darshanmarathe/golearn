package main

import (
	"fmt"
)


func main(){
	me := Person{"Darshan" , "Narayan" , "Marathe"}
	ShowFullNameWithMr(&me)
}


type Ifullname interface {
	fullname(title string)
}``

func ShowFullNameWithMr(r Ifullname){
	r.fullname("Mr")
}

type Person struct {
	FirstName string 
	MiddleName string 
	LastName string 
}


func(person *Person) fullname(title string){
	fmt.Println(title + " " +
person.FirstName  + " " +
person.MiddleName  + " " +
person.LastName)
}