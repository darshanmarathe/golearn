package main 

import "runtime"

func main(){
   runtime.GOMAXPROCS(10)

   ch := make(chan string)
   doneCh := make(chan bool)
   go abcgen(ch)
   go Printer(ch , doneCh)
  
  println("here it came first")
  <- doneCh
}


func abcgen(ch chan string){
	for i := byte('a'); i <= byte('z'); i++ {
	  ch <- string(i)
	}
	close(ch)
}


func Printer(ch chan string , doneCh chan bool)  {
	for i := range ch {
			println(i)
	}
	doneCh <- true
}