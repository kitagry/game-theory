package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
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

func NewRandom() *RandomPerson {
	return &RandomPerson{}
}

func (r *RandomPerson) Input(i Value) {
}

func (r *RandomPerson) Output() Value {
	rand.Seed(time.Now().UnixNano())
	return values[rand.Intn(len(values))]
}

type GrimTrigger struct {
	betrayed bool
}

func NewGrimTrigger() *GrimTrigger {
	return &GrimTrigger{betrayed: false}
}

func (g *GrimTrigger) Input(i Value) {
	if i == Betray {
		g.betrayed = true
	}
}

func (g *GrimTrigger) Output() Value {
	if g.betrayed {
		return Betray
	} else {
		return Trust
	}
}

type TipForTat struct {
	betrayed bool
}

func NewTipForTat() *TipForTat {
	return &TipForTat{betrayed: false}
}

func (t *TipForTat) Input(i Value) {
	if i == Betray {
		t.betrayed = true
	} else if i == Trust {
		t.betrayed = false
	}
}

func (t *TipForTat) Output() Value {
	if t.betrayed {
		return Betray
	} else {
		return Trust
	}
}

// PlayGame run game
func PlayGame(times int, p1, p2 Person) {
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

func GetPerson(i int) (Person, error) {
	if i == 0 {
		return NewRandom(), nil
	} else if i == 1 {
		return NewGrimTrigger(), nil
	} else if i == 2 {
		return NewTipForTat(), nil
	} else {
		return nil, errors.New("値が不正です.")
	}
}

func main() {
	times := flag.Int("n", 10, "Number of game attempts")
	p1Int := flag.Int("p1", 0, "Player1's Algorithm, 0: random, 1: Grim Trigger, 2: Tip for tat")
	p2Int := flag.Int("p2", 1, "Player2's Algorithm, 0: random, 1: Grim Trigger, 2: Tip for tat")
	flag.Parse()

	p1, err := GetPerson(*p1Int)
	if err != nil {
		log.Fatalln(err)
	}

	p2, err := GetPerson(*p2Int)
	if err != nil {
		log.Fatalln(err)
	}

	PlayGame(*times, p1, p2)
}
