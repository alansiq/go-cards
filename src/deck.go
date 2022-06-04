package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card

var cardSuits []string = []string{"Spades", "Hearts", "Diamonds", "Clovers"}
var cardValues []string = []string{
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
}

func newDeck() deck {
	createdDeck := deck{}
	for _, csuit := range cardSuits {
		for _, cvalue := range cardValues {
			card := card{suit: csuit, value: cvalue}
			createdDeck = append(createdDeck, card)
		}
	}
	return createdDeck
}

// go receiver
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card.toString())
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() (result string) {

	var cardsAsString []string

	for _, c := range d {
		cardsAsString = append(cardsAsString, c.toString())
	}

	result = strings.Join(cardsAsString, ",")

	return
}

func deckFromString(input []string) (result deck, err []string) {
	for _, c := range input {
		r, e := cardFromString(c)

		if len(e) == 0 {
			result = append(result, r)
			continue
		}
		err = append(err, e)

	}

	return
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("FAILED TO LOAD DECK FROM FILE: ", filename)
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	res, e := deckFromString(strings.Split(string(bs), ","))

	if len(e) > 0 {
		fmt.Println("Não foi possível converter os cards a seguir:")
		for i, problem := range e {
			fmt.Println(i, problem)
		}

	}

	return res
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}
