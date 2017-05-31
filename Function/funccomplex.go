package main 


import "fmt"

type name struct{
	firstname string 
	middlename string 
	lastname string
}


func main(){

    darshan := name{firstname : "darshan" , middlename:"narayan" , lastname: "Marathe"}
	PrintName( darshan )
}


func PrintName(pname name){
	fmt.Println(pname.firstname , pname.middlename , pname.lastname)
}