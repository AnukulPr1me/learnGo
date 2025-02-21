package main

import "fmt"

type Animal interface{
	hasLegs() string
	sound() string
}

type person struct {
	name string
	age int
}

func main(){
	firstPerson := person{name: "john", age: 30}
	fmt.Println(firstPerson.name)
	fmt.Println(firstPerson.age)
}