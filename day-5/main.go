package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (this_p Point) equals(that_p Point) bool {
	return this_p.X == that_p.X && this_p.Y == that_p.Y
}

func main() {
	inputFile := "input"
	fmt.Printf("%v\n", solvePart1(inputFile))
	fmt.Printf("%v\n", solvePart2(inputFile))
}

func parseFile(inputFile string) (lines []string) {
	f, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func parseLine(line string) (start Point, end Point) {
	tokens := strings.Split(line, " -> ")

	left := strings.Split(tokens[0], ",")
	leftX, _ := strconv.Atoi(string(left[0]))
	leftY, _ := strconv.Atoi(string(left[1]))

	right := strings.Split(tokens[1], ",")
	rightX, _ := strconv.Atoi(string(right[0]))
	rightY, _ := strconv.Atoi(string(right[1]))

	start = Point{leftX, leftY}
	end = Point{rightX, rightY}
	return
}

func pointsBetween(start Point, end Point) (points []Point) {
	points = append(points, start)

	p := start
	for !p.equals(end) {
		if p.X < end.X {
			p.X++
		}

		if p.X > end.X {
			p.X--
		}

		if p.Y < end.Y {
			p.Y++
		}

		if p.Y > end.Y {
			p.Y--
		}
		points = append(points, p)
	}

	return
}

func solvePart1(inputFile string) (answer int) {
	pointsCount := make(map[Point]int, 0)
	for _, line := range parseFile(inputFile) {
		start, end := parseLine(line)
		if !(start.X == end.X || start.Y == end.Y) {
			continue
		}

		for _, point := range pointsBetween(start, end) {
			pointsCount[point]++
		}
	}

	for _, count := range pointsCount {
		if count >= 2 {
			answer++
		}
	}
	return
}

func solvePart2(inputFile string) (answer int) {
	pointsCount := make(map[Point]int, 0)
	for _, line := range parseFile(inputFile) {
		start, end := parseLine(line)
		for _, point := range pointsBetween(start, end) {
			pointsCount[point]++
		}
	}

	for _, count := range pointsCount {
		if count >= 2 {
			answer++
		}
	}
	return
}
