package main

import (
	"fmt"
)

const (
	Empty = iota
	BlackStone
	WhiteStone
)

func main() {
	boardSize := 9
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}
	whiteKills := 0
	blackKills := 0

	minLetter := 'a'
	maxLetter := rune(97 + boardSize - 1)
	minNumber := 1
	maxNumber := boardSize
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

	move := 0
	for true {
		move++
		drawBoard(boardSize, board)
		fmt.Printf("White kills: %d, Black kills: %d\n", whiteKills, blackKills)
		fmt.Printf("Your move (%c%d-%c%d): ", minLetter, minNumber, maxLetter, maxNumber)
		var rowLetter rune
		var inputCol int
		// read both the rowLetter and the inputCol together as "a1" or "d5"
		fmt.Scanf("%c%d", &rowLetter, &inputCol)
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		row := int(rowLetter) - 97 // a=0, b=1, etc
		col := inputCol - 1        // convert 1-indexed input to 0-indexed
		if row < 0 || row >= boardSize || col < 0 || col >= boardSize {
			fmt.Println("Invalid move (out of bounds)")
			move--
			continue
		}
		if board[row][col] != Empty {
			fmt.Println("Invalid move (already occupied)")
			move--
			continue
		}
		color := 0
		if move%2 == 0 {
			color = WhiteStone
		} else {
			color = BlackStone
		}
		checked := make([][]bool, boardSize)
		for i := range checked {
			checked[i] = make([]bool, boardSize)
		}
		board[row][col] = color
		kills := killNeighbors(row, col, board, color)
		if color == WhiteStone {
			whiteKills += kills
		} else {
			blackKills += kills
		}
		libs := liberties(row, col, board, color)
		if libs == 0 {
			fmt.Println("Invalid move (no liberties)")
			move--
			board[row][col] = Empty
			continue
		}
		if color == WhiteStone {
			fmt.Printf("White played: %c%d\n", rowLetter, inputCol)
		} else {
			fmt.Printf("Black played: %c%d\n", rowLetter, inputCol)
		}
	}
}

func liberties(row int, col int, board [][]int, color int) int {
	boardSize := len(board)
	checked := make([][]bool, boardSize)
	for i := range checked {
		checked[i] = make([]bool, boardSize)
	}
	return libertiesRecursive(row, col, board, checked, color)
}

func libertiesRecursive(row int, col int, board [][]int, checked [][]bool, color int) int {
	if isOutOfBounds(row, col, board) || checked[row][col] {
		return 0
	}

	checked[row][col] = true

	if board[row][col] == Empty {
		return 1
	}

	if board[row][col] != color {
		return 0
	}

	count := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, direction := range directions {
		count += libertiesRecursive(row+direction[0], col+direction[1], board, checked, color)
	}
	return count
}

func isOutOfBounds(row int, col int, board [][]int) bool {
	return row < 0 || row >= len(board) || col < 0 || col >= len(board[row])
}

func killNeighbors(row int, col int, board [][]int, color int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	kills := 0
	for _, direction := range directions {
		r := row + direction[0]
		c := col + direction[1]
		if isOutOfBounds(r, c, board) || board[r][c] == Empty || board[r][c] == color {
			continue
		}
		if liberties(r, c, board, otherColor(color)) == 0 {
			kills += killGroup(r, c, board, otherColor(color))
		}
	}
	return kills
}
func killGroup(row int, col int, board [][]int, color int) int {
	if isOutOfBounds(row, col, board) || board[row][col] != color {
		return 0
	}
	board[row][col] = Empty
	kills := 1
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, direction := range directions {
		kills += killGroup(row+direction[0], col+direction[1], board, color)
	}
	return kills
}

func otherColor(color int) int {
	switch color {
	case WhiteStone:
		return BlackStone
	case BlackStone:
		return WhiteStone
	default:
		return Empty
	}
}
