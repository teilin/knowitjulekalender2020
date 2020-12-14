package main

import (
	"fmt"
	"math/big"
)

func generateSekvens(length int) ([]int, int) {
	var sekvens []int = make([]int, length)
	var seqContains map[int]int = make(map[int]int)
	count := 0

	sekvens[0] = 0
	seqContains[0]++
	sekvens[1] = 1
	seqContains[1]++

	for index := 2; index < length; index++ {
		tmp := sekvens[index-2] - index
		if tmp < 0 || seqContains[tmp] > 0 {
			tmp = sekvens[index-2] + index
		}
		i := big.NewInt(int64(tmp))
		if i.ProbablyPrime(1) {
			count++
		}
		seqContains[tmp]++
		sekvens[index] = tmp
	}
	return sekvens, count
}

func main() {
	_, count := generateSekvens(1800813)
	fmt.Println(count)
}
