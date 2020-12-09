package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Stack type
type Stack []Coor

// Coor type
type Coor struct {
	x int
	y int
}

// IsEmpty check for stack
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push to stack
func (s *Stack) Push(coor Coor) {
	*s = append(*s, coor)
}

// Pop value from stack
func (s *Stack) Pop() (Coor, bool) {
	if s.IsEmpty() {
		return Coor{}, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func readInput2(inoutFile string) map[Coor]rune {
	file, err := os.Open(inoutFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var states map[Coor]rune = make(map[Coor]rune)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		xs := []rune(line)
		for x, state := range xs {
			states[Coor{x: x, y: y}] = state
		}
		y++
	}
	return states
}

func checkNeighboursIsSick(states map[Coor]rune, coor Coor) bool {
	numSick := 0

	left := states[Coor{x: coor.x - 1, y: coor.y}]
	right := states[Coor{x: coor.x + 1, y: coor.y}]
	below := states[Coor{x: coor.x, y: coor.y - 1}]
	above := states[Coor{x: coor.x, y: coor.y + 1}]

	if left == 'S' {
		numSick++
	}
	if right == 'S' {
		numSick++
	}
	if below == 'S' {
		numSick++
	}
	if above == 'S' {
		numSick++
	}

	return numSick > 1
}

func traverse(states map[Coor]rune, numIterations *int) {
	*numIterations++
	var stack Stack
	shouldContinue := false
	for position, state := range states {
		if state == 'F' && checkNeighboursIsSick(states, position) {
			stack.Push(position)
		}
	}
	for stack.IsEmpty() == false {
		shouldContinue = true
		position, isPopped := stack.Pop()
		if isPopped {
			states[position] = 'S'
		}
	}
	if shouldContinue {
		traverse(states, numIterations)
	}
}

func main() {
	numIterations := 0
	states := readInput2("elves.txt")
	traverse(states, &numIterations)
	fmt.Println(numIterations)
}
