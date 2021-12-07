package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := loadLines("input")
	solvePowerConsumption(lines)
	solveLifeSupport(lines)
}

func loadLines(fname string) (lines []string) {
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func sumColumns(lines []string) (sums []int) {
	sums = make([]int, 0)
	for _, line := range lines {
		for i, r := range line {
			if len(sums) <= i {
				sums = append(sums, 0)
			}

			asNumber, _ := strconv.Atoi(string(r))
			sums[i] += asNumber
		}
	}
	return sums
}

func solvePowerConsumption(lines []string) {
	counts := sumColumns(lines)
	gammaBits := ""
	epsilonBits := ""
	threshold := len(lines) / 2

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

func solveLifeSupport(lines []string) {
	O2 := getO2(lines)
	Co2 := getCo2(lines)

	fmt.Printf("%d (o2) * %d (co2) = %d\n", O2, Co2, O2*Co2)
}

func getO2(lines []string) (O2 uint64) {
	for i := 0; i < len(lines[0]); i++ {
		var saved []string
		for _, line := range lines {
			if string(line[i]) == mostCommonAtIndex(lines, i) {
				saved = append(saved, line)
			}
		}

		lines = saved
		if len(lines) == 1 {
			O2Line := lines[0]
			O2, _ = strconv.ParseUint(string(O2Line), 2, len(O2Line))
			break
		}
	}

	return
}

func getCo2(lines []string) (Co2 uint64) {
	for i := 0; i < len(lines[0]); i++ {
		var saved []string
		for _, line := range lines {
			if string(line[i]) == leastCommonAtIndex(lines, i) {
				saved = append(saved, line)
			}
		}

		lines = saved
		if len(lines) == 1 {
			Co2Line := lines[0]
			Co2, _ = strconv.ParseUint(string(Co2Line), 2, len(Co2Line))
			break
		}
	}

	return
}

func leastCommonAtIndex(lines []string, i int) (s string) {
	s = mostCommonAtIndex(lines, i)
	if s == "1" {
		return "0"
	}
	return "1"
}

func mostCommonAtIndex(lines []string, i int) string {
	count := 0
	for _, line := range lines {
		if line[i] == '1' {
			count++
		}
		if line[i] == '0' {
			count--
		}
	}
	if count < 0 {
		return "0"
	}
	return "1"
}
