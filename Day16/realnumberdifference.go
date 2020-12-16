package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func abundantNumber(num int) (bool, int) {
	sum := 0
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			if num/i == i {
				sum += i
			} else {
				sum += i
				sum += (num / i)
			}
		}
	}
	sum -= num
	return sum > num, sum + num
}

func perfectSquare(num int) bool {
	i := int(math.Floor(math.Sqrt(float64(num))))
	return i*i == num
}

func main() {
	inputNumber, _ := strconv.Atoi(os.Args[1])
	start := time.Now()
	count := 0
	for i := 1; i < inputNumber; i++ {
		isAbundant, sum := abundantNumber(i)
		if isAbundant {
			diff := sum - (2 * i)
			if diff > 0 && perfectSquare(diff) {
				count++
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println(count)
	fmt.Printf("Execution time %s", elapsed)
}
