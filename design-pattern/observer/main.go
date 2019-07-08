package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Observer is
type Observer interface {
	Update(Subject)
}

// Subject is
type Subject interface {
	AddObserver(Observer)
	DeleteObserver(Observer)
	NotifyObservers()
	GetNumber() int
	Execute()
}

// NumberGenerator is
type NumberGenerator struct {
	observers []Observer
	number    int
}

// AddObserver returns
func (ng *NumberGenerator) AddObserver(o Observer) {
	ng.observers = append(ng.observers, o)
}

// DeleteObserver returns
func (ng *NumberGenerator) DeleteObserver(o Observer) {
	result := NumberGenerator{}
	for _, observer := range ng.observers {
		if observer != o {
			result.observers = append(result.observers, o)
		}
	}
	ng.observers = result.observers
}

// NotifyObservers returns
func (ng *NumberGenerator) NotifyObservers() {
	for _, observer := range ng.observers {
		observer.Update(ng)
	}
}

// GetNumber returns
func (ng *NumberGenerator) GetNumber() int {
	return ng.number
}

// Execute returns
func (ng *NumberGenerator) Execute() {
	rand.Seed(time.Now().UnixNano())

	// NOTE: randomな時間、numberを更新する
	for index := 0; index < 5; index++ {
		ng.number = rand.Intn(50)
		ng.NotifyObservers()

		st := rand.Intn(10)
		time.Sleep(time.Duration(st) * time.Second)
	}
}

// DegitObserver is
type DegitObserver struct{}

// Update returns
func (d *DegitObserver) Update(s Subject) {
	fmt.Println("DegitObserverが状態が変化したことを検知したよ: ", s.GetNumber())
}

func main() {
	ng := &NumberGenerator{}
	do := &DegitObserver{}
	ng.AddObserver(do)
	ng.Execute()
}
