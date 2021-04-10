package fizzbuzz_test

import (
	"testing"

	"github.com/anubanqwer/Go/fizzbuzz"
)

func TestFizzBuzzGiven1(t *testing.T) {
	given := 1
	want := "1"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven3(t *testing.T) {
	given := 3
	want := "Fizz"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven5(t *testing.T) {
	given := 5
	want := "Buzz"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven15(t *testing.T) {
	given := 15
	want := "FizzBuzz"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

//go test -v
//Run specific test case -> go test -run TestFizzBuzzGiven15
