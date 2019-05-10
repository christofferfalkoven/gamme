package main

import "time"
import "github.com/uu-os-2018/project-gamma/simulate"

//TODO: Write testcases for Board and Human
func main() {
	board := new(simulate.Board)
	board.AllocateBoard(100)
	time.Sleep(time.Second)
	simulate.StartServer(board)

}
