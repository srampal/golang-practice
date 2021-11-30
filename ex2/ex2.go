package ex2

import (
	"errors"

	"github.com/sirupsen/logrus"
)

//Exercise Fibonacci calculator. Takes user input n, returns Fib(n)
func Exercise(n uint32) (fib uint32, err error) {

	var i, fibMinus1, fibMinus2 uint32

	if n > 100 {
		logrus.Printf("Received invalid input %d\n", n)
		err = errors.New("Invalid input!")
	}

	if n == 0 {
		fib = 0
		return
	}
	if n == 1 {
		fib = 1
		return
	}

	fibMinus2 = 0
	fibMinus1 = 1

	for i = 2; i <= n; i++ {
		fib = fibMinus2 + fibMinus1
		fibMinus2 = fibMinus1
		fibMinus1 = fib
	}
	return
}
