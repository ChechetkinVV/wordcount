// wordcount tool counts words in a phrase.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	phrase := readInput()
	n := count(phrase)
	fmt.Println(n)
}

// count returns words amount in the phrase.
func count(phrase string) int {
	if phrase == "" {
		return 0
	}
	s := strings.Split(phrase, " ")
	return len(s)
}

// readInput reads phrase from command line and returns it.
func readInput() (str string) {
	str = os.Args[1]
	return str
}
