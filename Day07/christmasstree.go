package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Define const values
const (
	tree       = rune('#')
	blank rune = 32
)

func readFile(inputFile string) [][]rune {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var forrest [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		forrest = append(forrest, []rune(line))
	}
	return forrest
}

func isTreeSymmetric(forrest [][]rune, startX int, startY int) bool {
	isSymmetric := true

	for i := startY; i >= 0; i-- {
		j := 0
		for (-1*j)+startX > 0 && (forrest[i][(-1*j)+startX] != blank || forrest[i][j+startX] != blank) {
			//t1 := forrest[i][(-1*j)+startX]
			//fmt.Println(t1)
			//t2 := forrest[i][j+startX]
			//fmt.Println(t2)
			if (-1*j)+startX >= 0 && forrest[i][(-1*j)+startX] != forrest[i][j+startX] {
				return false
			}
			j++
		}
	}

	return isSymmetric
}

func main() {
	forrest := readFile("input.txt")
	forrestBottom := len(forrest) - 1
	numSymmetricTrees := 0

	for index, value := range forrest[forrestBottom] {
		if value == tree {
			if isTreeSymmetric(forrest, index, forrestBottom) {
				numSymmetricTrees++
			}
		}
	}

	fmt.Println(numSymmetricTrees)
}
