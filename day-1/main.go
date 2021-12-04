package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveIncreasing() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	timesIncreased := -1
	prev := -1
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		if depth > prev {
			timesIncreased++
		}
		prev = depth
	}

	println(timesIncreased)
}

func solveIncreasingWindow() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	depths := []int{}
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, depth)
	}

	timesIncreased := -1
	prev := -1
	for i := 0; i < len(depths)-2; i++ {
		window := depths[i : i+3]
		sum := 0
		for j := 0; j < len(window); j++ {
			sum += window[j]
		}

		if sum > prev {
			timesIncreased++
			fmt.Printf("%d (%d) > %d	total: %d\n", window, sum, prev, timesIncreased)
		} else {
			fmt.Printf("%d (%d) <= %d\n", window, sum, prev)
		}
		prev = sum
	}
	println(timesIncreased)
}

func main() {
	solveIncreasing()
	solveIncreasingWindow()
}
