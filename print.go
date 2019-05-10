package main

import "fmt"

func printBoard(board [][]*Human) {
	for row := range board {
		for col := range board[row] {
			fmt.Print(board[row][col].pos)
		}
		fmt.Println("")
	}
}

func printNeighbors(neighbors [][]*Human) {
	for row := range neighbors {
		fmt.Println("")
		for col := range neighbors[row] {
			fmt.Print(neighbors[row][col].pos)
		}
	}
	fmt.Println("")
}
