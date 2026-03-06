package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type German struct {
}

type Italian struct {
}

type Portuguese struct {
}

func (g German) LanguageName() string {
	return "German"
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (g German) Greet(n string) string {
	return fmt.Sprintf("Hallo %s", n)
}

func (i Italian) Greet(n string) string {
	return fmt.Sprintf("Ciao %s", n)
}

func (p Portuguese) Greet(n string) string {
	return fmt.Sprintf("Olá %s", n)
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", greeter.LanguageName(), greeter.Greet(name))
}
