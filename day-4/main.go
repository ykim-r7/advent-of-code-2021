package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input"
	fmt.Printf("winBingo final score: %d\n", winBingo(inputFile))
	fmt.Printf("loseBingo final score: %d\n", loseBingo(inputFile))
}

type Board struct {
	sequences [][]int
}

func removeIfContains(slice []int, this int) (newSlice []int) {
	newSlice = slice
	for idx, element := range slice {
		if element == this {
			newSlice = append(slice[:idx], slice[idx+1:]...)
		}
	}
	return
}

func (b *Board) selectNumber(selectedNumber int) {
	for i := 0; i < len(b.sequences); i++ {
		for j := 0; j < len(b.sequences[i]); j++ {
			if selectedNumber == b.sequences[i][j] {
				b.sequences[i] = removeIfContains(b.sequences[i], selectedNumber)
			}
		}
	}
}

func (b *Board) hasBingo() (bingo bool) {
	for i := 0; i < len(b.sequences); i++ {
		if len(b.sequences[i]) == 0 {
			bingo = true
		}
	}
	return
}

func (b *Board) calculateScore() (score int) {
	seen := make(map[int]bool)
	for i := 0; i < len(b.sequences); i++ {
		for j := 0; j < len(b.sequences[i]); j++ {
			remainingNumber := b.sequences[i][j]
			if !seen[remainingNumber] {
				seen[remainingNumber] = true
			}
		}
	}

	for k := range seen {
		score += k
	}

	return
}

func createBoard(numbers []int) (b Board) {
	b = Board{
		sequences: [][]int{},
	}
	n := int(math.Sqrt(float64(len(numbers))))
	if n == 0 {
		return
	}

	var sequence []int
	for i := 0; i < n; i++ { // ↓
		sequence = []int{}
		for j := 0; j < n; j++ { // →
			sequence = append(sequence, numbers[(i*n)+j]) // ↔
		}
		b.sequences = append(b.sequences, sequence)
	}

	for j := 0; j < n; j++ { // →
		sequence = []int{}
		for i := 0; i < n; i++ { // ↓
			sequence = append(sequence, numbers[(i*n)+j]) // ↕
		}
		b.sequences = append(b.sequences, sequence)
	}

	return
}

func parseFile(inputFile string) (calledNumbers []int, boards []Board) {
	f, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	firstLine := scanner.Text()
	slice := strings.Split(firstLine, ",")

	for _, s := range slice {
		num, _ := strconv.Atoi(s)
		calledNumbers = append(calledNumbers, num)
	}

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}

		var boardNumbers []int
		for i := 0; i < 5; i++ {
			strings := strings.Fields(t)
			for _, str := range strings {
				number, _ := strconv.Atoi(str)
				boardNumbers = append(boardNumbers, number)
			}
			scanner.Scan()
			t = scanner.Text()
		}

		board := createBoard(boardNumbers)
		boards = append(boards, board)
	}
	return
}

func winBingo(inputFile string) (score int) {
	calledNumbers, boards := parseFile(inputFile)
	for _, calledNumber := range calledNumbers {
		for _, board := range boards {
			board.selectNumber(calledNumber)
			if board.hasBingo() {
				score = board.calculateScore() * calledNumber
				return
			}
		}
	}
	return
}

func loseBingo(inputFile string) (score int) {
	calledNumbers, boards := parseFile(inputFile)

	winningBoardIndexes := make(map[int]bool, 0)
	for _, calledNumber := range calledNumbers {
		for i, board := range boards {
			if winningBoardIndexes[i] {
				continue
			}

			board.selectNumber(calledNumber)
			if board.hasBingo() {
				winningBoardIndexes[i] = true
				score = board.calculateScore() * calledNumber
			}
		}
	}
	return
}
