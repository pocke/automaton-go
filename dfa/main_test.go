package main

import (
	"bytes"
	"reflect"
	"testing"
)

var M1 = &Automaton{
	States:   []string{"q1", "q2", "q3"},
	Alphabet: []string{"0", "1"},
	Transitions: map[string]map[string]string{
		"q1": {
			"0": "q1",
			"1": "q2",
		},
		"q2": {
			"0": "q3",
			"1": "q2",
		},
		"q3": {
			"0": "q2",
			"1": "q2",
		},
	},
	StartState:   "q1",
	AcceptStates: []string{"q2"},
}

func TestNewAutomaton(t *testing.T) {
	r := bytes.NewReader([]byte(`q1 q2 q3
0 1
q1,0,q1 q1,1,q2 q2,0,q3 q2,1,q2 q3,0,q2 q3,1,q2
q1
q2`))
	a, err := NewAutomaton(r)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(a, M1) {
		t.Errorf("%+v != %+v", a, M1)
	}
}

func TestNewAutomatonWhenInvalidTransitionFunc(t *testing.T) {
	r := bytes.NewReader([]byte(`q1 q2 q3
0 1
q1,0,q1 q1,1,q2 q2,0,q3 q2,1,q2 q3,0 1,q2 q3,1,q2
q1
q2`))
	_, err := NewAutomaton(r)
	if err == nil {
		t.Error("should return error, but got nil")
	}
}

func TestRunAutomaton(t *testing.T) {
	assert := func(input []string, expected bool) {
		res, err := M1.Run(input)
		if err != nil {
			t.Error(err)
		}
		if res != expected {
			t.Errorf("Expected %v, but got %v", expected, res)
		}
	}

	assert([]string{}, false)
	assert([]string{"0", "1"}, true)
	assert([]string{"0", "1", "1"}, true)
	assert([]string{"0", "1", "0"}, false)
	assert([]string{"0", "1", "0", "0"}, true)
}
