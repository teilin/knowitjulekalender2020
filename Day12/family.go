package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readInput(inputFile string) string {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func main() {
	familyTreeStr := readInput(os.Args[1])
	familyTree := []rune(familyTreeStr)
	var generations map[int]int = make(map[int]int)
	generationCounter := 0
	prevChar := '#'
	for _, char := range familyTree {
		if char == ' ' {
			if prevChar != ')' {
				generations[generationCounter]++
			}
		} else if char == '(' {
			generationCounter++
		} else if char == ')' {
			if prevChar != ')' {
				generations[generationCounter]++
			}
			generationCounter--
		}
		prevChar = char
	}
	largestGeneration := 0
	for _, g := range generations {
		if largestGeneration < g {
			largestGeneration = g
		}
	}
	fmt.Println(largestGeneration)
}
