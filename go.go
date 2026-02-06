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
	boardSize := 19
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}

	move := 0
	for true {
		move++
		drawBoard(boardSize, board)
		fmt.Print("Your move (a1-s19): ")
		var rowLetter rune
		var inputCol int
		// read both the rowLetter and the inputCol together as "a1" or "d5"
		fmt.Scanf("%c%d", &rowLetter, &inputCol)
		row := int(rowLetter) - 97 // a=0, b=1, etc
		col := inputCol - 1        // convert 1-indexed input to 0-indexed
		if row < 0 || row >= boardSize || col < 0 || col >= boardSize || board[row][col] != Empty {
			fmt.Println("Invalid move")
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
		libs := liberties(row, col, board, checked, color)
		if libs == 0 {
			fmt.Println("Invalid move (no liberties)")
			move--
			board[row][col] = Empty
			continue
		}
	}
}

func liberties(row int, col int, board [][]int, checked [][]bool, color int) int {
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
		count += liberties(row+direction[0], col+direction[1], board, checked, color)
	}
	return count
}

func isOutOfBounds(row int, col int, board [][]int) bool {
	return row < 0 || row >= len(board) || col < 0 || col >= len(board[row])
}
