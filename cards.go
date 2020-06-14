package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type card struct {
	suit  string
	value string
}

func generateCard(suit string, value string) card {
	c := card{suit: suit}
	c.value = value
	return c
}

func cardvalue() string {
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

func suitvalue() (s string) {
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

func generateHand() {
	hand := make([]card, 2)
	hand[0] = generateCard(suitvalue(), cardvalue())
	time.Sleep(100 * time.Millisecond)
	hand[1] = generateCard(suitvalue(), cardvalue())
	fmt.Printf("Your hand is: %v, %v", hand[0], hand[1])
}

func main() {
	generateHand()
}
