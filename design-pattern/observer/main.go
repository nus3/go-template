package main

import (
	"fmt"
	"math/rand"
)

// Observer is
type Observer interface {
	Update()
}

// Generator is
type Generator interface {
	AddObserver(Observer)
	// TODO: 削除
	// DeleteObserver(Observer)
	NotifyObservers(Observer)
	GetNumber() int
	Execute()
}

// NumberGenerator is
type NumberGenerator struct {
	observers []*Observer
	number    int
}

// AddObserver returns
func (n *NumberGenerator) AddObserver(o *Observer) {
	n.observers = append(n.observers, o)
}

// TODO: observer

// NotifyObservers returns
func (n *NumberGenerator) NotifyObservers() {
	for _, observer := range n.observers {
		observer.Update()
	}
}

// GetNumber returns
func (n *NumberGenerator) GetNumber() int {
	return n.number
}

// Execute returns
func (n *NumberGenerator) Execute() {
	for index := 0; index < 20; index++ {
		n.number = rand.Intn(15)
		n.NotifyObservers()
	}
}

// DegitObserver is
type DegitObserver struct{}

// Update returns
func (d DegitObserver) Update(n *NumberGenerator) {
	fmt.Println(n.GetNumber())
}

// GraphObserver is
type GraphObserver struct{}

// Update returns
func (gr GraphObserver) Update(n *NumberGenerator) {
	count := n.GetNumber()
	output := ""

	for index := 0; index < count; index++ {
		output = output + "*"
	}

	fmt.Println(output)
}

func main() {
	generator := &NumberGenerator{}
	var d Observer = &DegitObserver{}
	var g Observer = &GraphObserver{}

	generator.AddObserver(d)
	generator.AddObserver(g)
	generator.Execute()
}
