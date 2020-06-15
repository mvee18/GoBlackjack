package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var ErrGenerateCard = errors.New("Could not generate the card.")

type card struct {
	suit  string
	value string
}

var AllCards = make([]card,0)

func generateCard(suit string, value string) (c card) {
	c = card{suit: suit}
	c.value = value

	_, b := SeeIfCardGenerated(AllCards, c)
	switch b {
	case true:
		c = generateCard(suitValue(), cardValue())
		return c
	}
	AllCards = append(AllCards, c)
	return c
}

func cardValue() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	value := r1.Intn(13)
	f, err := convertValue(value)
	if err != nil {
		log.Fatal("Could not convert the value.")
	}
	return f
}

func convertValue(i int) (paint string, err error) {
	switch i {
	case 0:
		paint := "A"
		return paint, err
	case 11:
		paint := "J"
		return paint, err
	case 12:
		paint := "Q"
		return paint, err
	case 13:
		paint := "K"
		return paint, err
	default:
		paint := strconv.Itoa(i)
		return paint, err
	}
}

func suitValue() (s string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	suit := r1.Intn(3)
	switch suit {
	case 0:
		s := "♠"
		return s
	case 1:
		s := "♥"
		return s
	case 2:
		s := "♦"
		return s
	case 3:
		s := "♣"
		return s
	default:
		s := "0"
		log.Fatal("Did not get a suit.")
		return s
	}
}

func generateHand() (hand []card, err error) {
	hand = make([]card, 2)
	hand[0] = generateCard(suitValue(), cardValue())
	time.Sleep(100 * time.Millisecond)
	hand[1] = generateCard(suitValue(), cardValue())
	fmt.Printf("Your hand is: %v, %v", hand[0], hand[1])
	return hand, err
}



func SeeIfCardGenerated (s []card, c card) (int, bool) {
	for i, item := range s {
		if item == c {
			return i, true
		}
	}
	return -1, false
}

func main() {
	generateHand()
}
