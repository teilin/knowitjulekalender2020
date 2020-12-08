package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// coor type
type coor struct {
	x int
	y int
}

// city type
type city struct {
	name     string
	position coor
	time     float64
}

func distanceCalculation(distanceTo *coor, currentCoor *coor) float64 {
	return math.Abs((float64(distanceTo.x) - float64(currentCoor.x))) + math.Abs((float64(distanceTo.y) - float64(currentCoor.y)))
}

func readInput(cities *[]city, route *[]string, inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ": ") {
			citySplit := strings.Split(line, ": ")
			var cityVar city
			cityVar.name = citySplit[0]
			coorSplit := strings.Split(citySplit[1], ", ")
			var coordinate coor
			coordinate.x, _ = strconv.Atoi(coorSplit[0][1:len(coorSplit[0])])
			coordinate.y, _ = strconv.Atoi(coorSplit[1][0 : len(coorSplit[1])-1])
			cityVar.position = coordinate
			*cities = append(*cities, cityVar)
		} else {
			*route = append(*route, line)
		}
	}
}

func move(coordinate *coor, x int, y int) {
	(*coordinate).x += x
	(*coordinate).y += y
}

func findDestCity(cities *[]city, cityName string) (int, city) {
	for index, c := range *cities {
		if c.name == cityName {
			return index, c
		}
	}
	return -1, city{name: "UKNOWN"}
}

func getTime(distance float64) float64 {
	if distance < 50.0 && distance >= 20.0 {
		return 0.75
	} else if distance < 20.0 && distance >= 5.0 {
		return 0.5
	} else if distance < 5.0 && distance > 0.0 {
		return 0.25
	} else if distance == 0.0 {
		return 0.0
	} else {
		return 1.0
	}
}

func updateTime(cities *[]city, currentPosition *coor) {
	for index, c := range *cities {
		distance := distanceCalculation(&c.position, currentPosition)
		(*cities)[index].time += getTime(distance)
	}
}

func travel(route *[]string, cities *[]city) {
	currentPosition := coor{x: 0, y: 0}
	for _, dest := range *route {
		_, c := findDestCity(cities, dest)
		for currentPosition.x < c.position.x || currentPosition.x > c.position.x {
			if currentPosition.x < c.position.x {
				move(&currentPosition, 1, 0)
			} else {
				move(&currentPosition, -1, 0)
			}
			updateTime(cities, &currentPosition)
		}
		for currentPosition.y < c.position.y || currentPosition.y > c.position.y {
			if currentPosition.y < c.position.y {
				move(&currentPosition, 0, 1)
			} else {
				move(&currentPosition, 0, -1)
			}
			updateTime(cities, &currentPosition)
		}
	}
}

func getMinMax(cities *[]city) (float64, float64) {
	sort.SliceStable(*cities, func(i, j int) bool {
		return (*cities)[i].time < (*cities)[j].time
	})
	return (*cities)[0].time, (*cities)[len(*cities)-1].time
}

func main() {
	var cities []city
	var route []string
	readInput(&cities, &route, "input.txt")
	travel(&route, &cities)
	min, max := getMinMax(&cities)
	diffMinMax := max - min
	fmt.Println(diffMinMax)
}
