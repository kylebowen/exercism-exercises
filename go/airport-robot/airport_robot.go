package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	lang := greeter.LanguageName()
	greeting := greeter.Greet(name)
	return fmt.Sprintf("I can speak %s: %s", lang, greeting)
}

type Italian struct{}

func (g Italian) LanguageName() string {
	return "Italian"
}

func (g Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (g Portuguese) LanguageName() string {
	return "Portuguese"
}

func (g Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
