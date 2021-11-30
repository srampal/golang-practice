package main

import (
	"fmt"

	"github.com/srampal/golang-practice/ex1"
	"github.com/srampal/golang-practice/ex2"
)

func main() {
	ex1.Exercise()

	ex2TestInput := []uint32{2, 3, 4, 5, 6, 200}

	for _, val := range ex2TestInput {
		ans, err := ex2.Exercise(val)
		if err == nil {
			fmt.Printf("ex2.Exercise(%d) returned %d \n", val, ans)
		} else {
			fmt.Print(err)
		}
	}

}
