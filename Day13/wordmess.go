package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	alphabeth = [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

func countChars(slice []rune, char rune) int {
	counter := 0
	for _, c := range slice {
		if c == char {
			counter++
		}
	}
	return counter
}

func remove(slice []rune, index int) []rune {
	return append(slice[:index], slice[index+1:]...)
}

func find(char rune) (int, rune) {
	for index, c := range alphabeth {
		if c == char {
			return index + 1, c
		}
	}
	panic("Unknown char")
}

func readInput(inputFile string) string {
	fileContent, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(fileContent)
}

func main() {
	inputs := []rune(readInput(os.Args[1]))
	var taken map[rune]int = make(map[rune]int)
	for i := 0; i < len(inputs); i++ {
		if taken[inputs[i]] == 0 && inputs[i] != ' ' {
			n, char := find(inputs[i])
			num := 0
			taken[char]++
			numOccuranses := countChars(inputs, inputs[i])
			if n <= numOccuranses {
				for j := 0; j < len(inputs); j++ {
					if inputs[j] == char {
						num = num + 1
						if num != n {
							inputs[j] = ' '
						}
					}
				}
			} else {
				for j := 0; j < len(inputs); j++ {
					if inputs[j] == char {
						inputs[j] = ' '
					}
				}
			}
		}
	}
	for a := 0; a < len(inputs); a++ {
		if inputs[a] == ' ' {
			inputs = remove(inputs, a)
		}
	}
	fmt.Println(strings.ReplaceAll(strings.TrimSpace(string(inputs)), " ", ""))
}
