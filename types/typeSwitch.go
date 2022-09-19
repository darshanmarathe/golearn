package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("its a bool \n")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Man struct {
	Name string
}

func main() {
	do(21)
	do("hello")
	do(true)

	raghu := &Man{Name: "Raghu"}
	do(raghu)

}
