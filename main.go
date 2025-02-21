package main

type Animal interface {
	hasLegs() string
	sound() string
}

type Dog struct {
	name  string
	legs  int
	sound string
}
type person struct {
	name  string
	legs  int
	sound string
}

func main() {
	firstDog := Dog{name: "Tommy", legs: 4, sound: "woof-woof"}
	firstPerson := person{name: "Anukul", legs: 2, sound: "language"}

}
