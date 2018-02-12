/*
Package prose is used to perform language-processing by the use of entity-matching.

GitHub: http://github.com/snwfdhmp/prose
Developer: snwfdhmp
*/
package prose

import (
	"io"

	log "github.com/sirupsen/logrus"
)

// AI represents a Bot
type AI struct {
	Name    string      //Name of the AI. Will be used for example in conversations
	Actions []*Action   //Actions the AI can handle
	Logger  *log.Logger //Logger to log to
}

// NewAI creates an AI with name and logger in params, and will initialize Actions to an empty []*Action
func NewAI(name string, logger *log.Logger) *AI {
	return &AI{name, make([]*Action, 0), logger}
}

// Handle will add actions in params to actions that AI can handle
func (ai *AI) Handle(actions ...*Action) {
	for _, a := range actions {
		ai.Actions = append(ai.Actions, a)
	}
}

// Process will return a slice of every Action needing to be run according to input and action triggers
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

// Run will run a slice of Action and output to w, it will stop at the first error if stopIfError=true
func (ai *AI) Run(actions []*Action, w io.Writer, stopIfError bool) []error {
	var errs []error
	for _, a := range actions {
		a.Writer = w
		if err := a.Run(a); err != nil {
			errs = append(errs, err)
			if stopIfError {
				return errs
			}
		}
	}

	return errs
}
