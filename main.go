package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
╱╲╳╭╮╯╰
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	w := []rune("╱╲")
	for r := 0; r < 24; r++ {
		for c := 0; c < 40; c++ {
			fmt.Printf("%c", w[rand.Intn(len(w))])
		}
		fmt.Println()
	}
}
