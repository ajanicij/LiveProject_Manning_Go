package main

import (
	"fmt"
	"time"
)

// The board dimensions.
const numRows = 8
const numCols = numRows

// Whether we want an open or closed tour.
const requireClosedTour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
	dr, dc int
}

var moveOffsets []Offset

var numCalls int64

var numVisitedMax int

func initializeOffsets() {
    moveOffsets = []Offset {
        Offset {-2, -1},
        Offset {-1, -2},
        Offset {+2, -1},
        Offset {+1, -2},
        Offset {-2, +1},
        Offset {-1, +2},
        Offset {+2, +1},
        Offset {+1, +2},
    }
}

func makeBoard(numRows, numCols int) [][]int {
	var board [][]int
	for row := 0; row < numRows; row++ {
		var board_row []int
		for col := 0; col < numCols; col++ {
			board_row = append(board_row, unvisited)
		}
		board = append(board, board_row)
	}
	return board
}

func dumpBoard(board [][]int) {
	for row := 0; row < len(board); row++ {
		boardRow := board[row]
		rowStr := ""
		for col := 0; col < len(boardRow); col++ {
			rowStr += fmt.Sprintf("%02d ", boardRow[col])
		}
		fmt.Println(rowStr)
	}
	fmt.Println()
}

func onBoard(numRows, numCols, tryRow, tryCol int) bool {
	if tryRow < 0 || tryCol < 0 {
		return false
	}
	if tryRow >= numRows || tryCol >= numCols {
		return false
	}
	return true
}

// Try to extend a knight's tour starting at (curRow, curCol).
// Return true or false to indicate whether we have found a solution.
func findTour(board [][]int, numRows, numCols, curRow, curCol, numVisited int) bool {
	// fmt.Printf("findTour: curRow=%d curCol=%d\n", curRow, curCol)
	// time.Sleep(10 * time.Millisecond)
	numCalls++
	if numCalls % 10000000 == 0 {
		// fmt.Printf("findTour: numCalls=%d numVisited=%d\n", numCalls, numVisited)
	}
	if numVisited > numVisitedMax {
		fmt.Printf("findTour: numCalls=%d numVisited=%d\n", numCalls, numVisited)
		fmt.Printf("findTour: numVisited=%d\n", numVisited)
		numVisitedMax = numVisited
	}
	// Check if the knight has visited every square.
	if numVisited == numRows * numCols {
		if !requireClosedTour {
			return true
		}
		// assert requireClosedTour == true
		for _, offset := range moveOffsets {
			r := curRow + offset.dr
			c := curCol + offset.dc
			if onBoard(numRows, numCols, r, c) && board[r][c] == 0 {
				return true
			}
		}
		return false
	}
	
	// The knight has not visited every square.
	for _, offset := range moveOffsets {
		r := curRow + offset.dr
		c := curCol + offset.dc
		if !onBoard(numRows, numCols, r, c) || board[r][c] != unvisited {
			continue
		}
		board[r][c] = numVisited
		if findTour(board, numRows, numCols, r, c, numVisited + 1) {
			// Found a tour.
			return true
		}
		board[r][c] = unvisited
	}
	// fmt.Printf("findTour backtrack: curRow=%d curCol=%d\n", curRow, curCol)
	return false
}

func main() {
    numCalls = 0

    // Initialize the move offsets.
    initializeOffsets()

    // Create the blank board.
    board := makeBoard(numRows, numCols)

    // Try to find a tour.
    start := time.Now()
    board[0][0] = 0
    if findTour(board, numRows, numCols, 0, 0, 1) {
        fmt.Println("Success!")
    } else {
        fmt.Println("Could not find a tour.")
    }
    elapsed := time.Since(start)
    dumpBoard(board)
    fmt.Printf("%f seconds\n", elapsed.Seconds())
    fmt.Printf("%d calls\n", numCalls)
}

