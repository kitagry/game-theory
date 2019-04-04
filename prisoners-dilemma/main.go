package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Value int

const (
	Betray Value = iota
	Trust
)

var (
	values []Value = []Value{Betray, Trust}

	// 1つ目のValueは自分, 2つ目のValueは相手
	Scores map[Value]map[Value]int = map[Value]map[Value]int{
		Betray: {Betray: 1, Trust: 5},
		Trust:  {Betray: 0, Trust: 4},
	}
)

type Person interface {
	Input(i Value)
	Output() Value
}

type RandomPerson struct {
}

func (r RandomPerson) Input(i Value) {
}

func (r RandomPerson) Output() Value {
	rand.Seed(time.Now().UnixNano())
	return values[rand.Intn(len(values))]
}

func Game(times int, p1, p2 Person) {
	p1Score := 0
	p2Score := 0
	for i := 0; i < times; i++ {
		p1Answer := p1.Output()
		p2Answer := p2.Output()

		p1Score += Scores[p1Answer][p2Answer]
		p2Score += Scores[p2Answer][p1Answer]

		p1.Input(p2Answer)
		p2.Input(p1Answer)
	}

	fmt.Printf("Player1's score is %d\n", p1Score)
	fmt.Printf("Player2's score is %d\n", p2Score)
}

func main() {
	p1 := RandomPerson{}
	p2 := RandomPerson{}

	Game(10, p1, p2)
}
