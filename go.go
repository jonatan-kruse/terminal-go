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
	// Create 0-indexed board matrix
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
		if move%2 == 0 {
			board[row][col] = WhiteStone
		} else {
			board[row][col] = BlackStone
		}
	}
}
