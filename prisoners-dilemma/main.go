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
	GetAlgorithm() string

	Input(i Value)
	Output() Value
}

type Betrayer struct {
}

func NewBetrayer() *Betrayer {
	return &Betrayer{}
}

func (b *Betrayer) GetAlgorithm() string {
	return "Betray"
}

func (b *Betrayer) Input(i Value) {
}

func (b *Betrayer) Output() Value {
	return Betray
}

type Truster struct {
}

func NewTruster() *Truster {
	return &Truster{}
}

func (t *Truster) GetAlgorithm() string {
	return "Trust"
}

func (t *Truster) Input(i Value) {
}

func (t *Truster) Output() Value {
	return Trust
}

type RandomPerson struct {
}

func NewRandom() *RandomPerson {
	return &RandomPerson{}
}

func (r *RandomPerson) GetAlgorithm() string {
	return "Random"
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

func (g *GrimTrigger) GetAlgorithm() string {
	return "Grim Trigger"
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

func (t *TipForTat) GetAlgorithm() string {
	return "Tip for Tat"
}

func (t *TipForTat) Input(i Value) {
	switch i {
	case Betray:
		t.betrayed = true
	case Trust:
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
func PlayGame(times int, p1, p2 Person) (int, int) {
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

	return p1Score, p2Score
}

func GetPerson(i int) (Person, error) {
	if i == 0 {
		return NewBetrayer(), nil
	} else if i == 1 {
		return NewTruster(), nil
	} else if i == 2 {
		return NewRandom(), nil
	} else if i == 3 {
		return NewGrimTrigger(), nil
	} else if i == 4 {
		return NewTipForTat(), nil
	} else {
		return nil, errors.New("値が不正です.")
	}
}

func CompareAll() {
	const algorithmNum = 5
	resultTable := "| 自分のアルゴリズム \\ 相手のアルゴリズム"
	for i := 0; i < algorithmNum; i++ {
		p, err := GetPerson(i)
		if err != nil {
			return
		}
		resultTable += fmt.Sprintf("| %s ", p.GetAlgorithm())
	}
	resultTable += "|\n"

	resultTable += "|"
	for i := 0; i < algorithmNum; i++ {
		resultTable += ":--:|"
	}
	resultTable += "\n"

	for i := 0; i < algorithmNum; i++ {
		p1, err := GetPerson(i)
		if err != nil {
			return
		}

		resultTable += fmt.Sprintf("| %s | ", p1.GetAlgorithm())

		for j := 0; j < algorithmNum; j++ {
			p1, err := GetPerson(i)
			if err != nil {
				return
			}

			p2, err := GetPerson(j)
			if err != nil {
				return
			}
			p1Score, _ := PlayGame(10, p1, p2)
			resultTable += fmt.Sprint(p1Score) + " | "
		}
		resultTable += "\n"
	}
	fmt.Print(resultTable)
}

func main() {
	times := flag.Int("n", 10, "Number of game attempts")
	p1Int := flag.Int("p1", 0, "Player1's Algorithm, 0: betray, 1: trust, 2: random, 3: Grim Trigger, 4: Tip for tat")
	p2Int := flag.Int("p2", 1, "Player2's Algorithm, 0: betray, 1: trust, 2: random, 3: Grim Trigger, 4: Tip for tat")
	playSingle := flag.Bool("o", true, "Play single play or not")
	flag.Parse()

	if *playSingle {
		p1, err := GetPerson(*p1Int)
		if err != nil {
			log.Fatalln(err)
		}

		p2, err := GetPerson(*p2Int)
		if err != nil {
			log.Fatalln(err)
		}

		p1Score, p2Score := PlayGame(*times, p1, p2)
		fmt.Printf("Player1's (%s) score is %d\n", p1.GetAlgorithm(), p1Score)
		fmt.Printf("Player2's (%s) score is %d\n", p2.GetAlgorithm(), p2Score)
	} else {
		CompareAll()
	}
}
