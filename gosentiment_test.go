package main

import "testing"

func TestRateText(t *testing.T) {
	rating := RateText("Xero smells very bad")
	if rating > 0 {
		t.Error("Sentence is negative but was rated positive.")
	}

	rating = RateText("This is not good.")
	if rating > 0 {
		t.Error("Sentence is negative but was rated positive.")
	}

	rating = RateText("This is not not good.")
	if rating < 0 {
		t.Error("Sentence is positive but was rated negative.")
	}
}
