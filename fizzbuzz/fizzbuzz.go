package fizzbuzz

import "strconv"

func Say(n int) string {
	fizz := n%3 == 0
	buzz := n%5 == 0
	if fizz && buzz {
		return "FizzBuzz"
	} else if fizz {
		return "Fizz"
	} else if buzz {
		return "Buzz"
	} else {
		return strconv.Itoa(n)
	}
}
