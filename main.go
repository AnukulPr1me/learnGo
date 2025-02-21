package main

import "fmt"

type newThing interface{
	helloworld()
	newWorld()
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