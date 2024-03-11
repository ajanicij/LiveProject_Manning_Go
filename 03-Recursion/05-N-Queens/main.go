package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const numRows = 6
	board := makeBoard(numRows)

	start := time.Now()
	success := placeQueens1(board, numRows, 0, 0)
	//success := placeQueens2(board, numRows, 0, 0, 0)
	//success := placeQueens3(board, numRows, 0, 0, 0)

	elapsed := time.Since(start)
	if success {
		fmt.Println("Success!")
		dumpBoard(board, numRows)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}

func makeBoard(numRows int) [][]string {
	var board [][]string
	for r := 0; r < numRows; r++ {
		row := make([]string, numRows)
		for c := 0; c < numRows; c++ {
			row[c] = "."
		}
		board = append(board, row)
	}
	return board
}

func dumpBoard(board [][]string, numRows int) {
	for r := 0; r < numRows; r++ {
		row := board[r]
		fmt.Println(strings.Join(row, " "))
	}
}

// Return true if this series of squares contains at most one queen.
func seriesIsLegal(board [][]string, numRows, r0, c0, dr, dc int) bool {
	r := r0
	c := c0
	count := 0
	for {
		if r < 0 || r >= numRows {
			break
		}
		if c < 0 || c >= numRows {
			break
		}
		// fmt.Printf("seriesIsLegal: r=%d c=%d\n", r, c)
		if board[r][c] == "Q" {
			count++
		}
		r += dr
		c += dc
	}
	return count < 2
}

// Return true if the board is legal.
func boardIsLegal(board [][]string, numRows int) bool {
	// Check rows.
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, numRows, r, 0, 0, 1) {
			return false
		}
	}

	// Check columns.
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, numRows, 0, c, 1, 0) {
			return false
		}
	}
	// Check diagonals.
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, numRows, r, 0, 1, 1) {
			return false
		}
		if !seriesIsLegal(board, numRows, r, 0, -1, 1) {
			return false
		}
	}
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, numRows, 0, c, 1, 1) {
			return false
		}
		if !seriesIsLegal(board, numRows, numRows-1, c, -1, 1) {
			return false
		}
	}
	return true
}

// Return true if the board is legal and a solution.
func boardIsASolution(board [][]string) bool {
	count := 0
	numRows := len(board)
	if !boardIsLegal(board, numRows) {
		return false
	}
	for r := 0; r < numRows; r++ {
		for c := 0; c < numRows; c++ {
			if board[r][c] == "Q" {
				count++
			}
		}
	}
	return count == numRows
}

// Try placing a queen at position [r][c].
// Return true if we find a legal board.
func placeQueens1(board [][]string, numRows, r, c int) bool {
	if r >= numRows {
		return boardIsASolution(board)
	}
	nextR := r
	nextC := c + 1
	if nextC == numRows {
		nextR++
		nextC = 0
	}
	// 1. try not placing a queen in the square (r, c).
	if placeQueens1(board, numRows, nextR, nextC) {
		return true
	}
	// 2. try placing a queen in the square (r, c).
	board[r][c] = "Q"
	if placeQueens1(board, numRows, nextR, nextC) {
		return true
	}
	// Did not find a solution. Backtrack.
	board[r][c] = "."
	return false
}

