package main

import "fmt"

const (
	N = 9
)

func main() {
	board := [N][N]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}
	printBoard(board)
}

func printBoard(board [N][N]int) {
	for i, _ := range board {
		if i % 3 == 0 && i != 0 {
			fmt.Println(" _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _")
		}
		for j, _ := range board[i] {
			if j % 3 == 0 && j != 0 {
				fmt.Printf(" | ")
			}
			if j == 8 {
				fmt.Printf(" %v \n", board[i][j])
			} else {
				fmt.Printf(" %v ", board[i][j])
			}
		}
	}
}
