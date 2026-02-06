package main

import (
	"fmt"
)

func drawBoard(boardSize int, board [][]int) {
	// Wood-colored background
	bg := "\033[48;2;220;179;92m"
	fg := "\033[38;2;60;60;60m"
	border := "\033[48;2;200;160;70m"
	paddedBoardSize := boardSize + 2

	reset := "\033[0m"

	fmt.Print(border)
	for col := 0; col <= (paddedBoardSize+2)*2; col++ {
		fmt.Print(" ")
	}
	fmt.Print(reset)
	fmt.Println()
	for row := 0; row < paddedBoardSize; row++ {
		for col := 0; col < paddedBoardSize; col++ {
			if row == 0 || col == 0 || row == paddedBoardSize-1 || col == paddedBoardSize-1 {
				fmt.Print(border)
				fmt.Print(fg)
				if row == 0 && col != 0 && col != paddedBoardSize-1 {
					// the number padded with spaces to be in total 2 chars
					if col < 10 {
						fmt.Print(" ")
					}
					fmt.Print(col)
					if col == paddedBoardSize-2 {
						fmt.Print(" ")
					}
				} else if row != 0 && row != paddedBoardSize-1 && col == 0 {
					fmt.Print("  ")
					// The letter a-z
					fmt.Print(string(rune(97+row-1)) + " ")
				} else if row != paddedBoardSize-1 {
					fmt.Print("    ")
				}
				fmt.Print(reset)
				continue
			}
			fmt.Print(bg)
			fmt.Print(fg)
			if board[row-1][col-1] == WhiteStone {
				if col == 1 {
					fmt.Print(" ") // Center
				}
				fmt.Print("⚪")
			} else if board[row-1][col-1] == BlackStone {
				if col == 1 {
					fmt.Print(" ") // Center
				}
				fmt.Print("⚫")
			} else if isDot(row, col, paddedBoardSize) {
				fmt.Print("⏺─")
			} else if row == 1 && col == 1 {
				fmt.Print(" ┌─") // Top-left corner
			} else if row == 1 && col == boardSize {
				fmt.Print("┐ ") // Top-right corner
			} else if row == boardSize && col == 1 {
				fmt.Print(" └─") // Bottom-left corner
			} else if row == boardSize && col == boardSize {
				fmt.Print("┘ ") // Bottom-right corner
			} else if row == 1 {
				fmt.Print("┬─") // Top edge
			} else if row == boardSize {
				fmt.Print("┴─") // Bottom edge
			} else if col == 1 {
				fmt.Print(" ├─") // Left edge
			} else if col == boardSize {
				fmt.Print("┤ ") // Right edge
			} else {
				fmt.Print("┼─") // Center intersections
			}
			fmt.Print(reset)
		}
		if row != paddedBoardSize-1 {
			fmt.Println()
		}
	}
	fmt.Print(border)
	for col := 0; col <= (paddedBoardSize+2)*2; col++ {
		fmt.Print(" ")
	}
	fmt.Print(reset)
	fmt.Println()
	fmt.Print(border)
	for col := 0; col <= (paddedBoardSize+2)*2; col++ {
		fmt.Print(" ")
	}
	fmt.Print(reset)
	fmt.Println()
}

func isDot(row int, col int, boardSize int) bool {
	min := 4
	max := boardSize - 5
	if boardSize < 13 {
		min = 3
		max = boardSize - 4
	}
	if row == min && col == min {
		return true
	}
	if row == max && col == max {
		return true
	}
	if row == min && col == max {
		return true
	}
	if row == max && col == min {
		return true
	}
	if boardSize%2 == 1 && row == boardSize/2 && col == boardSize/2 {
		return true
	}

	return false
}
