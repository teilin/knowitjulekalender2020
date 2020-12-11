package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	alphabeth = [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

func indexLetter(char rune) int {
	for index, c := range alphabeth {
		if c == char {
			return index
		}
	}
	return -1
}

func addLetters(a rune, b rune) rune {
	ia := indexLetter(a)
	ib := indexLetter(b)
	r := (ia + ib) % 26
	return alphabeth[r]
}

func nextLetter(char rune) rune {
	index := 0
	for i, c := range alphabeth {
		if c == char {
			index = i
			break
		}
	}
	if index == len(alphabeth)-1 {
		return alphabeth[0]
	}
	return alphabeth[index+1]
}

func removeFirst(slice []rune) []rune {
	var tmp []rune = make([]rune, len(slice)-1)
	for i := 1; i < len(slice); i++ {
		tmp[i-1] = slice[i]
	}
	return tmp
}

func shiftLetters(slice []rune) []rune {
	for i, c := range slice {
		slice[i] = nextLetter(c)
	}
	return slice
}

func seekVertically(hints [][]rune, pass []rune, x int, y int) bool {
	ylength := len(hints)
	for _, p := range pass {
		xlength := len(hints[y])
		if x < xlength {
			if p != hints[y][x] {
				return false
			}
			y++
			if y >= ylength && p != pass[len(pass)-1] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func findHint(hintBreakdown [][]rune, password []rune) bool {
	for y, row := range hintBreakdown {
		for x, c := range row {
			if c == password[0] {
				if seekVertically(hintBreakdown, password, x, y) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	file, _ := os.Open("hint.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var matrix [][]rune
		originalString := scanner.Text()
		seekPass := "eamqia"
		if len(originalString) >= len(seekPass) {
			runes := []rune(originalString)
			matrix = append(matrix, runes)
			for i := 0; i < len(matrix[0]); i++ {
				runes = matrix[len(matrix)-1]
				runes = removeFirst(runes)
				runes = shiftLetters(runes)
				for j, c := range runes {
					runes[j] = addLetters(c, matrix[len(matrix)-1][j])
				}
				matrix = append(matrix, runes)
			}
			if findHint(matrix, []rune(seekPass)) {
				fmt.Println(originalString)
			}
		}
	}
}
