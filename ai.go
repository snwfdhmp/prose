/*
# Prose

## Introduction

Prose is a golang package used to perform language-processing by the use of entity-matching.

Extended documentation can be found [here](https://godoc.org/github.com/snwfdhmp/prose).

## Summary

- [Get Started](#get-started)
	- [Entity](#entity)
	- [Action](#action)
	- [AI](#ai)
- [Example](#example)
- [Maintainer](#example)

# Get Started

You can install package simply by

```
go get github.com/snwfdhmp/prose
```

Then you can use it simply by importing `github.com/snwfdhmp/prose`.

Full documentation can be found [here](https://godoc.org/github.com/snwfdhmp/prose).

## Entity

An entity object contains a compiled regexp that will match strings if the entity exists in it.
You can consider an entity as a named-regexp.

Example:

An entity can check *words* ...

```go
what := prose.NewEntity("what", prose.RegexpWords("what"))
```

... it can also detect *topics* or *abstract concepts* ...
```go
sport := prose.NewEntity("sport", prose.RegexpWords("football", "basketball", "hockey"))
love := prose.NewEntity("love", prose.RegexpWords("love", "romance", "wife"))
```

... or even *character types*.
```go
number := prose.NewEntity("number", "[1-9]")
```

> All that you can track with regexps, you can track it with prose.

## Action

An action is _an object containing a function_ that will be run in a particular context.

Example :

An action can be a *shell command* ...

```go
ai := prose.NewAI("Jarvis", logrus.New())
action := ai.NewCommandAction("echo 'Ran at $(date +%H:%M:%S)'")
```

... or a *custom function* of your choice.

```go
ai := prose.NewAI("Jarvis", logrus.New())

run := func (a *prose.Action) error {
	ai.Logger.Println("I'm an awesome func.")

	io.WriteString(a.Writer, "And I can also write to streams.")
}

action := prose.NewAction(run)
```

## AI

An AI is an object containing the *executive part of the scheme*.
It will reference actions to run, and will process the input string from _(the user|an http request|a slack message|whatever)_

They are simply created by :

```go
ai := prose.NewAI("Jarvis", logrus.New())
```

# Example

```go
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
}
```

Let's run the example above :

```sh
$ go run example.go
It's 01:53
```

# Maintainer

My GitHub profile : [snwfdhmp](https://github.com/snwfdhmp)
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
