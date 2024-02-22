package main

import (
	"fmt"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
)

func main() {
	f, err := os.Create("tickets_6.prof")
	if err != nil {
		fmt.Println(err)
		return
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	numberCount := 6

	happyCount := 0
	maxN := int(math.Pow(10, float64(numberCount)) - 1)

	k := strconv.Itoa(numberCount)

	format := "%" + k + "d"

	for i := 0; i <= maxN; i++ {
		ticketNumber := fmt.Sprintf(format, i)

		if isHappy(ticketNumber) {
			happyCount += 1
		}
	}

	fmt.Println("Happy Count ", happyCount)
}

func isHappy(s string) bool {
	half := len(s) / 2

	firstPart := 0

	secondPart := 0

	for i := 0; i < half; i++ {
		j, _ := strconv.Atoi(s[i : i+1])
		firstPart += j
	}

	for j := half; j < len(s); j++ {
		i, _ := strconv.Atoi(s[j : j+1])
		secondPart += i
	}

	return firstPart == secondPart
}