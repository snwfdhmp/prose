# Prose

## Introduction

Prose is a golang package used to perform language-processing by the use of entity-matching.

It's developed by [snwfdhmp](https://github.com/snwfdhmp)

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