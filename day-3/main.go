package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	solvePowerConsumption()
	solveLifeSupport()
}

var threshold int
var counts []int
var lines [][]rune

func solvePowerConsumption() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	counts = make([]int, len(lines[0]))

	for _, buf := range lines {
		for i, r := range buf {
			asNumber, _ := strconv.Atoi(string(r))
			counts[i] += asNumber
		}
	}

	gammaBits := ""
	epsilonBits := ""
	threshold = len(lines) / 2
	for _, count := range counts {
		if count > threshold {
			gammaBits += "1"
			epsilonBits += "0"
		} else {
			gammaBits += "0"
			epsilonBits += "1"
		}
	}
	gamma, _ := strconv.ParseUint(gammaBits, 2, len(gammaBits))
	epsilon, _ := strconv.ParseUint(epsilonBits, 2, len(epsilonBits))
	fmt.Printf("Gamma (%d) * Epsilon (%d) = %d\n", gamma, epsilon, gamma*epsilon)
}

func solveLifeSupport() {
	var o2 uint64
	var co2 uint64

	var remainingLines [][]rune = lines
	for i := 0; i < len(counts); i++ {
		var savedLines [][]rune
		mostCommon := mostCommonAtIndex(remainingLines, i, '1')
		for _, line := range remainingLines {
			if line[i] == mostCommon {
				savedLines = append(savedLines, line)
			}
		}
		remainingLines = savedLines
		if len(remainingLines) == 1 {
			o2, _ = strconv.ParseUint(string(remainingLines[0]), 2, len(remainingLines[0]))
			break
		}
	}

	remainingLines = lines
	for i := 0; i < len(counts); i++ {
		var savedLines [][]rune
		leastCommon := leastCommonAtIndex(remainingLines, i, '0')
		for _, line := range remainingLines {
			if line[i] == leastCommon {
				savedLines = append(savedLines, line)
			}
		}
		remainingLines = savedLines
		if len(remainingLines) == 1 {
			co2, _ = strconv.ParseUint(string(remainingLines[0]), 2, len(remainingLines[0]))
			break
		}
	}
	fmt.Printf("%d (o2) * %d (co2) = %d\n", o2, co2, o2*co2)
}

func mostCommonAtIndex(lines [][]rune, i int, tiebreaker rune) rune {
	count := 0
	for _, line := range lines {
		if line[i] == '1' {
			count++
		}
		if line[i] == '0' {
			count--
		}
	}
	if count > 0 {
		return '1'
	}
	if count < 0 {
		return '0'
	}
	return tiebreaker
}

func leastCommonAtIndex(lines [][]rune, i int, tiebreaker rune) rune {
	count := 0
	for _, line := range lines {
		if line[i] == '1' {
			count++
		}
		if line[i] == '0' {
			count--
		}
	}
	if count > 0 {
		return '0'
	}
	if count < 0 {
		return '1'
	}
	return tiebreaker
}
