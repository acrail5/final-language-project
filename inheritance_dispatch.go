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
