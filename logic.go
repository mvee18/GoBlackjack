package main

import (
	"errors"
	"fmt"
)

var ErrPlayerHand = errors.New("could not generate player hand")

func PlayerHand() (hand []card, err error) {
	PlayerHand, err := generateHand()
	if err != nil {
		return PlayerHand, ErrPlayerHand
	}
	fmt.Printf("Your hand is %v, %v", PlayerHand[0], PlayerHand[1])
	return PlayerHand, err
}

func DealerHand() (hand []card, err error) {
	DealerHand, err := generateHand()
	if err != nil {
		return DealerHand, ErrPlayerHand
	}
	fmt.Printf("The Dealer's Hand is {X X}, %v", DealerHand[1])
	return DealerHand, err
}

func GameLogic() {
	PlayerHand()
	DealerHand()
	println("What will you do?")
}
