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

type Game struct {
	board [N][N]int
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

	//board2 := [N][N]int{
	//{5, 3, 0, 0, 7, 0, 0, 0, 0},
	//{6, 0, 0, 1, 9, 5, 0, 0, 0},
	//{0, 9, 8, 0, 0, 0, 0, 6, 0},
	//{8, 0, 0, 0, 6, 0, 0, 0, 3},
	//{4, 0, 0, 8, 0, 3, 0, 0, 1},
	//{7, 0, 0, 0, 2, 0, 0, 0, 6},
	//{0, 6, 0, 0, 0, 0, 2, 8, 0},
	//{0, 0, 0, 4, 1, 9, 0, 0, 5},
	//{0, 0, 0, 0, 8, 0, 0, 7, 9},
	//}

	printBoard(&board)
	solution, err := solve(&board)
	if err == nil {
		fmt.Println("Sudoku board solution 😎🚀🚀🚀")
		printBoard(solution)
	} else {
		panic(err)
	}
}

func solve(board *[N][N]int) (*[N][N]int, error) {
	if isGameOver(board) {
		return board, nil
	}

	cell, err := checkEmpty(board)
	if err != nil {
		return board, err
	}

	for val := 1; val < 10; val++ {
		board[cell.x][cell.y] = val
		if valid(board, val, cell) {
			sol, err := solve(board)
			if err == nil {
				return sol, nil
			}
		}
		board[cell.x][cell.y] = 0
	}

	return board, errors.New("Unable to solve game")
}

func solveBoard(board *[N][N]int) {
	for row := 1; row <= 9; row++ {
		for col := 1; col <= 9; col++ {
			if board[row-1][col-1] == 0 {
				for value := 1; value < 10; value++ {
					cell := &Pos{row - 1, col - 1}
					if valid(board, value, cell) {
						board[cell.x][cell.y] = value
						solveBoard(board)
						board[cell.x][cell.y] = 0
					}
				}
				return
			}
		}
	}
	printBoard(board)
	return
}

func valid(board *[N][N]int, num int, p *Pos) bool {
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

func checkEmpty(board *[N][N]int) (*Pos, error) {
	for i, row := range board {
		for j, val := range row {
			if val == 0 {
				return &Pos{i, j}, nil
			}
		}
	}
	return nil, errors.New("no empty case found")
}

func isGameOver(board *[N][N]int) bool {
	for _, line := range board {
		for _, val := range line {
			if val == 0 {
				return false
			}
		}
	}
	return true
}

func printBoard(board *[N][N]int) {
	for i, _ := range board {
		if i%3 == 0 && i != 0 {
			fmt.Println(" - - - - - - - - - - - - - - - -")
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
	fmt.Printf("\n\n")
}
