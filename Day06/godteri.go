package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(inputFile string) []int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var input []int
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for value := range record {
			intVal, _ := strconv.Atoi(record[value])
			input = append(input, intVal)
		}
		return input
	}
}

func sum(input []int, start int, end int) int {
	sum := 0
	for i := start; i < end; i++ {
		sum += input[i]
	}
	return sum
}

func main() {
	input := readFile("godteri.txt")
	numPackagesOpen := len(input)
	numAlves := 127
	numPiecesCandy := sum(input, 0, numPackagesOpen)
	for numPiecesCandy%numAlves != 0 {
		numPackagesOpen--
		numPiecesCandy = sum(input, 0, numPackagesOpen)
	}
	fmt.Println(strconv.Itoa(numPiecesCandy/numAlves) + "," + strconv.Itoa(numPackagesOpen))
}
