package prose

import (
	"io"
	"os/exec"
)

//Action represents an Action
type Action struct {
	Writer  io.Writer
	Trigger [][]*Entity
	Run     func(*Action) error
}

// NewAction returns a new action that runs the function in params if triggered
// The functions have to be func(*Action) error, in Action you can use Writer that
// can change through every execution of Run command if needed (for example for API,
// chat bot, SMTP, ...)
func NewAction(run func(*Action) error) *Action {
	return &Action{Run: run, Trigger: newTrigger()}
}

// On adds a trigger to the action, that will be triggered if every entity in params
// match the tested string
func (a *Action) On(entities ...*Entity) {
	a.Trigger = append(a.Trigger, entities)
}

func (a *Action) Write(str string) {
	io.WriteString(a.Writer, str)
}

// NewCommandAction is a shortcut used to provide shell-command-running actions
func (ai *AI) NewCommandAction(command string) *Action {
	return NewAction(ai.NewCommandActionFunc(command))
}

// NewCommandActionFunc is a shortcut used to render a Run function executing a shell command
func (ai *AI) NewCommandActionFunc(command string) func(a *Action) error {
	return func(a *Action) error {
		ai.Logger.Infoln("Running", command)
		cmd := exec.Command("zsh", "-c", command)

		return cmd.Run()
	}
}
