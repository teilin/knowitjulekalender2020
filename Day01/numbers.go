package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	var s []int
	for {
		record, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}
		for value := range record {
			intVal, _ := strconv.Atoi(record[value])
			s = append(s, intVal)
		}
		sort.Ints(s)
		for index, _ := range s {
			if index+1 < 99999 {
				if s[index]+1 != s[index+1] {
					fmt.Println(index + 2)
				}
			}
		}
	}
}
