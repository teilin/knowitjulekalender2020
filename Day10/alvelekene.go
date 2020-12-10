package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("leker.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rank map[string]int = make(map[string]int)
	for scanner.Scan() {
		gren := strings.Split(scanner.Text(), ",")
		for index, elv := range gren {
			rank[elv] += len(gren) - (index + 1)
		}
	}
	winnerElv := ""
	winnerElvPoint := 0
	for elv, points := range rank {
		if winnerElvPoint < points {
			winnerElv = elv
			winnerElvPoint = points
		}
	}
	fmt.Println(winnerElv + "-" + strconv.Itoa(winnerElvPoint))
}
