package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type State int

const (
	OffHook    State = 1
	Connecting State = 2
	Connected  State = 3
	OnHold     State = 4
	OnHook     State = 5
)

func (s State) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

type Trigger int

const (
	CallDialed    Trigger = 1
	HungUp        Trigger = 2
	CallConnected Trigger = 3
	PlacedOnHold  Trigger = 4
	TakenOffHold  Trigger = 5
	LeftMessage   Trigger = 6
)

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

type TriggerResult struct {
	Trigger Trigger
	State   State
}

var rules = map[State][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OnHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OnHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OnHook},
	},
}

func main() {
	state, exitState := OffHook, OnHook
	for ok := true; ok; ok = state != exitState {
		fmt.Printf("The phone is currently %s\n", state)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			tr := rules[state][i]
			fmt.Printf("%d. %s\n", i, tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := rules[state][i]
		state = tr.State
	}

	fmt.Printf("We are done using the phone\n")
}
