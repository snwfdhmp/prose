package prose

import (
	"io"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type AI struct {
	Name    string
	Actions []*Action
	Logger  *log.Logger
}

func NewAI(name string, logger *log.Logger) *AI {
	return &AI{name, make([]*Action, 0), logger}
}

func (ai *AI) Handle(actions ...*Action) {
	for _, a := range actions {
		ai.Actions = append(ai.Actions, a)
	}
}

func (ai *AI) Process(input string) []*Action {
	toExecute := make([]*Action, 0)
	for _, a := range ai.Actions {
		actionShouldBeExecuted := false
		for _, condition := range a.Trigger {
			conditionIsTrue := true
			for _, part := range condition {
				if !part.IsPresent.MatchString(input) {
					conditionIsTrue = false
					break
				}
			}
			if conditionIsTrue {
				actionShouldBeExecuted = true
				break
			}
		}
		if actionShouldBeExecuted {
			toExecute = append(toExecute, a)
		}
	}
	return toExecute
}

func (ai *AI) Run(actions []*Action, w io.Writer, stopIfError bool) []error {
	var errs []error
	for _, a := range actions {
		a.Writer = w
		ai.Logger.Infoln("Running", a.Command)
		cmd := exec.Command("zsh", "-c", a.Command)

		cmd.Stdout = a.Writer
		cmd.Stderr = a.Writer
		if err := cmd.Run(); err != nil {
			errs = append(errs, err)
			if stopIfError {
				return errs
			}
		}
	}

	return errs
}
