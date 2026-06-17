package main

import "fmt"

type Parent struct {
	Name string
}

func (p Parent) Speak() {
	fmt.Println("Parent Speak:", p.Name)
}

type Child struct {
	Parent
	Name string
}

func (c Child) Speak() {
	fmt.Println("Child Speak:", c.Name)
}

type Speaker interface {
	Speak()
}

func callSpeaker(s Speaker) {
	s.Speak()
}
func main() {
	child := Child{
		Parent: Parent{Name: "Parent Name"},
		Name:   "Child Name",
	}

	fmt.Println("Variable field test:")
	fmt.Println("child.Name =", child.Name)
	fmt.Println("child.Parent.Name =", child.Parent.Name)

	fmt.Println()

	fmt.Println("Method dispatch test:")
	child.Speak()
	child.Parent.Speak()

	fmt.Println()

	fmt.Println("Interface method dispatch:")
	callSpeaker(child)
}
