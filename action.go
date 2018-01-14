package prose

import (
	"io"
	"os/exec"
)

type Action struct {
	Writer  io.Writer
	Trigger [][]*Entity
	Run     func(*Action) error
}

func NewAction(run func(*Action) error) *Action {
	return &Action{Run: run, Trigger: NewTrigger()}
}

func NewTrigger() [][]*Entity {
	ts := make([][]*Entity, 0)
	for i, _ := range ts {
		ts[i] = make([]*Entity, 0)
	}
	return ts
}

func (a *Action) On(entities ...*Entity) {
	a.Trigger = append(a.Trigger, entities)
}

func (ai *AI) NewCommandAction(command string) *Action {
	return NewAction(ai.NewCommandActionFunc(command))
}

func (ai *AI) NewCommandActionFunc(command string) func(a *Action) error {
	return func(a *Action) error {
		ai.Logger.Infoln("Running", command)
		cmd := exec.Command("zsh", "-c", command)

		return cmd.Run()
	}
}
