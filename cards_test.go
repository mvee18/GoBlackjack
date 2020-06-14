package main

import "testing"

func TestgenerateCard(t *testing.T) {
	got := generateCard("Heart", "10")
	want := card{"Heart", "10"}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

