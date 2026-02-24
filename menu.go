package main

import "fmt"

func chooseBoardSize() int {
	sizes := map[string]int{"1": 9, "2": 13, "3": 19}
	for {
		fmt.Println()
		fmt.Println("Select board size:")
		fmt.Println("  1) 9x9")
		fmt.Println("  2) 13x13")
		fmt.Println("  3) 19x19")
		fmt.Print("Choose [1-3]: ")
		var input string
		fmt.Scanln(&input)
		if input == "" {
			return 19
		}
		if size, ok := sizes[input]; ok {
			return size
		}
		// Be kind to the user, if they enter "9" instead of "1" we can accept that :)
		if input == "9" {
			return 9
		}
		if input == "13" {
			return 13
		}
		if input == "19" {
			return 19
		}
		fmt.Println("\nInvalid choice, try again.")
	}
}
