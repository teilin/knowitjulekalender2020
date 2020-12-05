package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

type coordinate struct {
	x float64
	y float64
}

func readFile(inputFile string) []rune {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return []rune(string(content))
}

func calculateArea(moves []rune) float64 {
	coor1 := coordinate{x: 0, y: 0}
	coor2 := coordinate{x: 0, y: 0}
	var char rune = moves[0]
	var length float64 = 0
	var sum float64 = 0
	for index, c := range moves {
		if char != c || index == len(moves) {
			if char == rune('H') {
				coor2.x = coor1.x + length
				coor2.y = coor1.y
			} else if char == rune('V') {
				coor2.x = coor1.x - length
				coor2.y = coor1.y
			} else if char == rune('O') {
				coor2.x = coor1.x
				coor2.y = coor1.y + length
			} else if char == rune('N') {
				coor2.x = coor1.x
				coor2.y = coor1.y - length
			}
			sum += (coor1.x * coor2.y) - (coor1.y * coor2.x)
			coor1 = coor2
			char = c
			length = 0
		}
		length++
	}
	return math.Abs(sum / 2)
}

func main() {
	array := readFile("rule.txt")
	area := calculateArea(array)
	fmt.Println(area)
}
