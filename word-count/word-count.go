package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	splitted := strings.Fields(s)
	mapio := make(map[string]int)

	for i := range splitted {
		mapio[splitted[i]] = 0
	}

	for i := range splitted {
		count := 0
		for j := range splitted {
			if splitted[i] == splitted[j] {
				count += 1
			}
			mapio[splitted[i]] = count
		}
	}

	return mapio
}

func main() {
	wc.Test(WordCount)
}
