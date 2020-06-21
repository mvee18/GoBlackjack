package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateCard(t *testing.T) {
		t.Run("testing single card generation", func(t *testing.T) {
			got := generateCard("Heart", "10")
			want := card{"Heart", "10"}

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})

		t.Run("test if card was appended", func(t *testing.T) {
			got := generateCard("Club", "7")
			want := AllCards[1]

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("see if it can stop the same card from being generated", func(t *testing.T) {
			AllCards = []card{{"Heart", "10"}, {"Club", "6"}}
			c := generateCard("Heart", "10")
			card := card{"Heart", "10"}
			if c == card {
				t.Errorf("expected an error but did not get one")
			}

		})
}

func TestSeeIfCardGenerated(t *testing.T) {
	t.Run("testing with predetermined array and card", func(t *testing.T) {
		s := []card{{"Heart", "10"}, {"Club", "9"}}
		c := card{"Club", "9"}
		i, b := SeeIfCardGenerated(s, c)
		if b == true {
			fmt.Printf("Card is present at location %d", i)
		}
		if b == false {
			t.Errorf("Expected an error but did not get one.")
		}
	})
}

func TestDetermineValue(t *testing.T) {
	t.Run("testing without ace", func(t *testing.T) {
		got, _, _ := determineValue([]card{{"Club", "K"}, {"Heart", "6"}})
		want := 16

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("testing with ace", func(t *testing.T) {
		ace1, ace2, _ := determineValue([]card{{"Diamonds", "A"}, {"Club", "A"}})
		want1 := 2
		want2 := 22

		if ace1 != want1 {
			t.Errorf("got %d, wanted %d", ace1, want1)
		} else if ace2 != want2 {
			t.Errorf("got %d, want %d", ace2, want2)
		}
	})
	t.Run("testing if over 21", func(t *testing.T) {
		face1, face2, bust := determineValue([]card{{"Heart", "K"}, {"Club","K"},{"Diamonds", "K"}})
		fmt.Printf("%v", bust)
		if bust == nil {
			t.Errorf("Expected an error, but did not get one.")
		}
		fmt.Printf("Your total is %d or %d", face1, face2)
	})
}

func TestUserActions(t *testing.T) {
	newhand, total, ace, err  := UserActions([]card{{"Heart", "10"},{"Club","K"}})
	fmt.Printf("%v", newhand)
	fmt.Printf("%d or %d", total, ace)
	if err != nil {
		t.Errorf("got an error, but did not expect one.")
	}
}

func TestDealerLogic(t *testing.T) {
	dhand, dtotal, dace := DealerLogic([]card{{"Heart", "2"},{"Club", "3"}}, 5, 5, 10, 14)
	fmt.Printf("Hand: %v, %d, %d", dhand, dtotal, dace)
}

func TestPotMoney(t *testing.T) {
	value := PotMoney()
	want := Pot(15)
	if value != want {
		t.Errorf("want %v got %v", want, value)
	}
}