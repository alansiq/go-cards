package main

import (
	"strings"
)

type card struct {
	suit  string
	value string
}

func (c card) toString() string {
	return c.value + " of " + c.suit
}

func cardFromString(c string) (result card, err string) {
	// "Ace of Spades"

	split := strings.Split(c, " ")
	if len(split) < 3 {
		err = c
	} else {
		result = card{suit: split[2], value: split[0]}
	}

	return
}
