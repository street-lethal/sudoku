package main

import (
	"fmt"
	"sudoku/src"
)

func main() {
	board, err := src.LoadFromFile("./data/board.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	board.Init()

	src.Search(board)
	fmt.Println(board)

	fmt.Println(board.Validate())
}
