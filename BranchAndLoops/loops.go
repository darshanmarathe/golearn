package main

func main(){
	for_loops()
	foreach()
}


func foreach()  {
	s := []string {"Darshan" , "Rajesh" , "Amol"}
	for idx, v :=  range s {
		println(idx  , "=>" ,v)
	}



	fe := make(map[string]string)
	fe["boriwali"] ="nilakshi"
	fe["bhandup"] = "Nivedita"
	fe["pune"] = "Aditi"

	for k ,y := range fe {
		println(k ,y)
	}
}

func for_loops()  {
	
	for index := 0; index < 30; index++ {
		println(index)
	}

	i := 0
	for {
		i++ 
		println(i)
	if i > 5 {
		
	  break	
	}
	}

}