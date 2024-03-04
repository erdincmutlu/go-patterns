package main

import "fmt"

// light: on/off
type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct{}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning the light off...")
	sw.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning the light on...")
	sw.State = NewOnState()
}
func main() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
