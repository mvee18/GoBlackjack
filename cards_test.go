package main

import (
	"fmt"
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