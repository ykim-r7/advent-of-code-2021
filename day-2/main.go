package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveDive() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	x_axis := 0
	y_axis := 0
	aim := 0
	for scanner.Scan() {
		t := scanner.Text()
		split_t := strings.Split(t, " ")
		cmd := split_t[0]
		number, _ := strconv.Atoi(split_t[1])

		switch cmd {
		case "forward":
			{
				x_axis += number
				y_axis += aim * number
			}
		case "up":
			{
				aim -= number
			}
		case "down":
			{
				aim += number
			}
		}
		fmt.Printf("(%s) x-axis: %d, y-axis: %d\n", t, x_axis, y_axis)
	}

	fmt.Printf("SOLUTION: x-axis * y-axis = %d\n", x_axis*y_axis)
}

func main() {
	solveDive()
}
