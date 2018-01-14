package prose

import (
	"io"
)

type Action struct {
	Command string
	Writer  io.Writer
	Trigger [][]*Entity
}

func NewAction(command string) *Action {
	ts := make([][]*Entity, 0)
	for i, _ := range ts {
		ts[i] = make([]*Entity, 0)
	}
	return &Action{Command: command, Trigger: ts}
}

func (a *Action) On(entities ...*Entity) {
	a.Trigger = append(a.Trigger, entities)
}
