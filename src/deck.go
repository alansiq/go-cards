package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
	"time"
)

type deck []string

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
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// go receiver
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() (result string) {

	result = strings.Join([]string(d), ",")

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

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}

