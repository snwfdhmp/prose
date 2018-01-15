package main

import (
	"fmt"
	"os"
	gotime "time"

	"github.com/sirupsen/logrus"
	"github.com/snwfdhmp/prose"
)

var (
	log = logrus.New()
)

func main() {
	ai := prose.NewAI("Jarvis", log) // We create an AI named Jarvis

	// Let's create an entity matching 'what'
	what := prose.NewEntity("what", prose.RegexpWords("what"))
	// ...and another matching 'time'
	time := prose.NewEntity("time", prose.RegexpWords("time"))

	// Now we create an action : writing time
	sendTime := prose.NewAction(func(a *prose.Action) error {
		msg := fmt.Sprintf("It's %s\n", gotime.Now().Format("15:04"))
		a.Write(msg)
		return nil
	})
	sendTime.On(what, time) // it'll be run if 'what' and 'time' are in the tested strign
	ai.Handle(sendTime)     // it'll be run by ai

	input := "Hey ! What time is it ?" //Let's say that's our user input

	actions := ai.Process(input) // Our AI processes the input and decide which actions to run

	// Now we tell the AI to runs the actions, to output to w, and to stop at first error
	errs := ai.Run(actions, os.Stdout, true) // note that we receive a slice of errors
	if len(errs) > 0 {
		log.Errorln("Errors:", errs)
		return
	}

	// Output:
	// It's 01:53
}
