package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var ErrPlayerHand = errors.New("could not generate player hand")
var ErrDealerHand = errors.New("could not generate dealer hand")
var ErrHandBust = errors.New("hand value over 21\n")
var ErrDealerBust = errors.New("The dealer has busted. You win the pot.\n")
var ErrPlayerLoss = errors.New("The dealer wins. You lose your bet.\n")
var ErrPlayerOutOfMoney = errors.New("\nYou are out of money. Game over.\n")
var ErrNotEnoughMoney = errors.New("\nYou do not have enough money to double down.\n")

type Pot float64
type Money float64

var Bet Money = 0
var PotValue Money = 0
var PlayerMoney Money = 100

func PlayerHand() (hand []card, err error) {
	PlayerHand, err := generateHand()
	if err != nil {
		return PlayerHand, ErrPlayerHand
	}
	fmt.Printf("Your hand is %v, %v\n", PlayerHand[0], PlayerHand[1])
	return PlayerHand, err
}

func DealerHand() (hand []card, err error) {
	DealerHand, err := generateHand()
	if err != nil {
		return DealerHand, ErrPlayerHand
	}
	//	fmt.Printf("The Dealer's Hand is {X X}, %v\n", DealerHand[1])
	fmt.Printf("The dealer's hand is {X, X}, %v\n", DealerHand[1])
	return DealerHand, err
}

func determineValue(hand []card) (total int, ace int, bust error) {
	for _, c := range hand {
		val, b := paintConverter(c.value)
		if b == true {
			ace1, ace2 := AceConverter(c.value)
			total += ace1
			ace += ace2
		} else if b == false {
			total += val
			ace += val
		}
	}
	if total > 21 && ace > 21 {
		return total, ace, ErrHandBust
	} else if total > 21 && ace <= 21 {
		return 99, ace, bust
	} else {
		return total, ace, bust
	}
}

func AceConverter(a string) (int, int) {
	if a == "A" {
		return 1, 11
	}
	return 0, 0
}

func paintConverter(s string) (int, bool) {
	switch s {
	case "A":
		return 1, true
	case "K":
		return 10, false
	case "Q":
		return 10, false
	case "J":
		return 10, false
	default:
		i, _ := strconv.Atoi(s)
		return i, false
	}
}

// The player should only be able to affect their own hand.
func UserActions(playerhand []card) ([]card, int, int, error) {
	println("\nWhat will you do?")
	println("The options are: Hit, Stay, or Double Down (DD)")
	doubleDown := false

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("There was an error reading your input. Please try again.")
		_, _, _, _ = UserActions(playerhand)
	}
	input = strings.TrimSuffix(input, LineBreak)

	if input == "Hit" {
		newhand, total, ace, err := Hit(playerhand)
		if err != nil {
			fmt.Printf("Your hand has exceed 21.\n")
			PotResolution(ErrPlayerLoss)
		}
		fmt.Printf("Your new hand is %v\n", newhand)
		fmt.Printf("Your new total is %d or %d\n", total, ace)
		UserActions(newhand)
		return newhand, total, ace, nil

	} else if input == "Stay" {
		hand, total, ace, err := Stay(playerhand)
		if err != nil {
			fmt.Printf("%v", ErrHandBust)
		}
		return hand, total, ace, nil

	} else if (input == "Double Down" || input == "DD") && doubleDown == false {
		dderror := DoubleDown()
		if dderror != nil {
			fmt.Println(dderror)
			UserActions(playerhand)
		}
		doubleDown = true
		newhand, total, ace, err := Hit(playerhand)
		if err != nil {
			fmt.Printf("Your hand has exceed 21.\n")
			PotResolution(ErrPlayerLoss)
		}
		fmt.Printf("Your new hand is %v\n", newhand)
		fmt.Printf("Your new total is %d or %d\n", total, ace)
		UserActions(newhand)
		return newhand, total, ace, nil

	} else if playerhand[0].value == playerhand[1].value {
		Split(playerhand)
		
	} else {
		fmt.Printf("There was an error reading your input. Try again.")
		UserActions(playerhand)
	}
	return nil, 0, 0, nil

}

func DealerLogic(dHand []card, dTotal int, dAce int, pTotal int, pAce int) ([]card, int, int) {
	fmt.Printf("Dealer: %d, %d. Player: %d, %d\n", dTotal, dAce, pTotal, pAce)
	if dTotal >= pTotal && dAce >= pAce && (dAce <= 21 || dTotal <= 21) {
		fmt.Println("The dealer elects to stay.")
		fmt.Printf("The dealer's hand is %v\n", dHand)
		dhand, dtotal, dace, err := Stay(dHand)
		if err != nil {
			fmt.Printf("%v", ErrHandBust)
		}
		PotResolution(ErrPlayerLoss)
		return dhand, dtotal, dace

	} else {
		fmt.Println("The dealer elects to hit.")
		newdhand, dtotal, dace, err := Hit(dHand)
		time.Sleep(1 * time.Second)
		fmt.Printf("The dealer's new hand is %v\n", newdhand)
		if err != nil {
			PotResolution(ErrDealerBust)
		}
		newdhand, dtotal, dace = DealerLogic(newdhand, dtotal, dace, pTotal, pAce)
		return newdhand, dtotal, dace
	}
}

// TODO: Resolve pot loss/win and add/subtract the bet amount from the player's money.
func PotResolution(err error) {
	if err == ErrDealerBust {
		PlayerMoney += PotValue
		fmt.Printf("The dealer loses. You win the pot.")
		GameLogic()
	} else if err == ErrPlayerLoss {
		fmt.Printf("%v", ErrPlayerLoss)
		GameLogic()
	}
}

func PotMoney() Money {
	PotValue = 0
	if PlayerMoney <= 0 {
		log.Fatal(ErrPlayerOutOfMoney)
	}
	fmt.Printf("\nYour current money is %v\n", PlayerMoney)
	fmt.Printf("How much will you bet?\n")
	PotReader := bufio.NewReader(os.Stdin)
	input, err := PotReader.ReadString('\n')
	if err != nil {
		fmt.Println("There was an error reading your input. Please try again.")
	}
	input = strings.TrimSuffix(input, LineBreak)
	bet, _ := strconv.Atoi(input)
	Bet = Money(bet)
	if Bet <= PlayerMoney {
		PlayerMoney -= Bet
		PotValue += Bet * 3 / 2
		fmt.Printf("Your balance is now %v\n", PlayerMoney)
		fmt.Printf("The pot value is now %v\n", PotValue)
		return PotValue
	} else if Bet > PlayerMoney {
		println("You cannot bet more money than you have! Try again.")
		PotMoney()
		return PotValue
	} else if Bet == 0 {
		println("You must be some amount of money to play. Try again.")
		PotMoney()
		return PotValue
	}
	return PotValue
}

func Split(c []card) {
	if c[0].value == c[1].value {
		fmt.Println("You also have the option to split.")
	}
}

func DoubleDown() error {
	if PlayerMoney < PotValue * 2/3 {
		return ErrNotEnoughMoney
	}
	PlayerMoney -= PotValue * 2/3
	PotValue = PotValue * 2
	fmt.Printf("Your new balance is %v\n", PlayerMoney)
	fmt.Printf("The pot is now %v\n", PotValue)
	return nil
}

func GameLogic() {
	PotMoney()

	phand, err := PlayerHand()
	if err != nil {
		log.Fatalf("%v", ErrPlayerHand)
	}
	total, ace, bust := determineValue(phand)
	if bust != nil {
		log.Fatal("Error in initial hand generation.")
	}
	fmt.Printf("Your total is %v, or %v\n", total, ace)

	dhand, err := DealerHand()
	if err != nil {
		log.Fatalf("%v", ErrDealerHand)
	}
	dTotal, dAce, err := determineValue(dhand)

	_, PTotal, PAce, UserErr := UserActions(phand)
	if UserErr != nil {
		fmt.Println("There was an error reading your input, try again.")
	}

	NewDHand, NewDTotal, NewDAce := DealerLogic(dhand, dTotal, dAce, PTotal, PAce)
	fmt.Printf("The dealer's hand is %v, with a total of %d or %d\n", NewDHand, NewDTotal, NewDAce)
}
