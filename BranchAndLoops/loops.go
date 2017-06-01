package main

import (
	"fmt"
	"time"
)

func main() {
	Basic_For_Loop()
	WhileLopp()
	for_with_break()
	foreach()
	infinit()
}


func Basic_For_Loop(){
	fmt.Println("Basic for loop")
	for index := 0; index < 5; index++ {
		println(index)
	}
	fmt.Println("Done")
	fmt.Println("")
	fmt.Println("")

}


func WhileLopp(){
	fmt.Println("While for loop")
	index := 0
	for  index < 5  {
		index++
		println(index)
	}
	fmt.Println("Done")
	fmt.Println("")
	fmt.Println("")
}

func for_with_break() {

	i := 0
	for {
		i++
		println(i)
		if i > 5 {

			break
		}
	}

}


func foreach() {
	s := []string{"Darshan", "Rajesh", "Amol"}
	for idx, v := range s {
		println(idx, "=>", v)
	}

	fe := make(map[string]string)
	fe["boriwali"] = "nilakshi"
	fe["bhandup"] = "Nivedita"
	fe["pune"] = "Aditi"

	for k, y := range fe {
		println(k, y)
	}
}

func infinit(){
	i := 0;
	for {
		i++
		if i % 2 ==0 {
			continue
		}
		//gap := i * 100
		time.Sleep(1000)
		fmt.Println("Printing till infinity => " , i )
	}
}