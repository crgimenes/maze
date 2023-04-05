package main

import (
	"math/rand"
	"time"
)

/*
╱╲╳╭╮╯╰
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	w := []rune("╱╲")
	l := len(w)
	for {
		print(string(w[rand.Intn(l)]))
	}
}
