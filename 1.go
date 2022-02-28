package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name    string
	Country string
	Samurai bool
}

func (h *Human) String() string {
	if h.Samurai {
		return fmt.Sprintf("%s samurai from %s", h.Name, h.Country)
	} else {
		return fmt.Sprintf("%s from %s is not samurai", h.Name, h.Country)
	}
}

func (h *Human) Genpuku(samurai bool) {
	h.Samurai = samurai
}

type Action struct {
	Human
}

func NewAction(name string, Country string, Samurai bool) *Action {
	Country = strings.Title(strings.ToLower(Country))
	name = strings.Title(strings.ToLower(name))
	return &Action{
		Human{
			Name:    name,
			Country: Country,
			Samurai: Samurai,
		},
	}
}

func main() {
	action := NewAction("Nosir", "russia", false)

	fmt.Println(action)

	action.Genpuku(true)

	fmt.Println(action)
}
