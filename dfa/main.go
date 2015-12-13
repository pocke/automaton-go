package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Automaton struct {
	States       []string
	Alphabet     []string
	Transitions  map[string]map[string]string // [State][Alphabet]State
	StartState   string
	AcceptStates []string
}

func main() {
	res, err := Main(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func Main(r io.Reader) (bool, error) {
	a, err := NewAutomaton(r)
	if err != nil {
		return false, err
	}

	sc := bufio.NewScanner(r)
	sc.Scan()
	input := strings.Split(sc.Text(), " ")

	res, err := a.Run(input)
	if err != nil {
		return false, err
	}
	return res, nil
}

func NewAutomaton(r io.Reader) (*Automaton, error) {
	a := &Automaton{}
	sc := bufio.NewScanner(r)

	// States
	sc.Scan()
	a.States = strings.Split(sc.Text(), " ")

	// Alphabet
	sc.Scan()
	a.Alphabet = strings.Split(sc.Text(), " ")

	// Transition function
	a.Transitions = make(map[string]map[string]string)
	sc.Scan()
	for _, v := range strings.Split(sc.Text(), " ") {
		t := strings.Split(v, ",")
		// TODO: check len
		q := t[0]
		s := t[1]
		q2 := t[2]
		if _, exist := a.Transitions[q]; !exist {
			a.Transitions[q] = make(map[string]string)
		}
		a.Transitions[q][s] = q2
	}

	// Start State
	sc.Scan()
	a.StartState = sc.Text()

	// Set of Accept States
	sc.Scan()
	a.AcceptStates = strings.Split(sc.Text(), " ")

	// TODO: validate
	return a, nil
}

func (a *Automaton) Run(input []string) (bool, error) {
	// TODO: input check

	q := a.StartState
	for _, i := range input {
		q = a.Transitions[q][i]
	}

	for _, ac := range a.AcceptStates {
		if ac == q {
			return true, nil
		}
	}
	return false, nil
}
