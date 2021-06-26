package main

import (
	"errors"
	"fmt"
)

const (
	N = 9
)

type Pos struct {
	x, y int
}

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

func valid(board [N][N]int, num int, p *Pos) bool {
	boxI := p.x / 3
	boxJ := p.y / 3

	// check row
	for i, _ := range board {
		if board[p.x][i] == num && p.y != i {
			return false
		}
	}

	// check col
	for i, _ := range board {
		if board[i][p.y] == num && p.x != i {
			return false
		}
	}

	// check box
	for i := boxI * 3; i < boxI+3; i++ {
		for j := boxJ * 3; j < boxJ+3; j++ {
			if board[i][j] == num && p.x != i && p.y != j {
				return false
			}
		}
	}

	return true
}

func checkEmpty(board [N][N]int) (*Pos, error) {
	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == 0 {
				return &Pos {i, j}, nil
			}
		}
	}
	return nil, errors.New("No empty field")
}

func printBoard(board [N][N]int) {
	for i, _ := range board {
		if i%3 == 0 && i != 0 {
			fmt.Println(" _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _")
		}
		for j, _ := range board[i] {
			if j%3 == 0 && j != 0 {
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
