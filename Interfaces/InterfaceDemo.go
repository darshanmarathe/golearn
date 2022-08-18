package main

import (
	"fmt"
)

func main() {
	me := Person{"Darshan", "Narayan", "Marathe"}
	ironMan := &SuperHero{NickName: "Iron Man"}
	Tiger := Animal{"Tiger"}

	fmt.Println(me.SayHello())
	fmt.Println(ironMan.SayHello())
	fmt.Println(Tiger.SayHello())

	me.MakeMeSuperHero(ironMan)
	fmt.Print(ironMan)
	// fmt.Print(&me)

}

func ShowFullNameWithMr(r Ifullname) {
	r.fullname("Mr")
}

type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
}

type SuperHero struct {
	NickName   string
	FirstName  string
	MiddleName string
	LastName   string
	Powers     string
}

type Animal struct {
	Name string
}

type Ifullname interface {
	fullname(title string)
}

type IGreet interface {
	SayHello() string
}

type ISuperHeroMaker interface {
	MakeMeSuperHero(per *SuperHero)
}

func (person *Person) SayHello() string {
	return "Hello " + person.FirstName + " " + person.LastName
}

func (person *Person) MakeMeSuperHero(sup *SuperHero) {
	sup.FirstName = person.FirstName
	sup.LastName = person.LastName
	sup.MiddleName = person.MiddleName

}

func (animal *Animal) SayHello() string {
	if animal.Name == "cat" {
		return "myau	"
	}
	if animal.Name == "dog" {
		return "Bhow bhow	"
	}
	return "grrrrrrr"
}

func (hero *SuperHero) SayHello() string {
	return "Hello " + hero.NickName
}

func (person *Person) fullname(title string) {
	fmt.Println(title + " " +
		person.FirstName + " " +
		person.MiddleName + " " +
		person.LastName)
}
