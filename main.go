package main

import (
	"math/rand"
	"os"
)

/*
╱╲╳╭╮╯╰
*/

func main() {
	w := []rune("╱╲")
	l := len(w)
	for {
		os.Stdout.WriteString(string(w[rand.Intn(l)]))
	}
}
