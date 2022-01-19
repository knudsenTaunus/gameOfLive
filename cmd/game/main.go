package main

import (
	"fmt"
	"time"
)

const (
	WIDTH  = 8
	HEIGHT = 5
)

func main() {
	board := NewBoard()

	/*board[0][1].live = true
	board[0][5].live = true
	board[1][3].live = true
	board[2][1].live = true
	board[3][3].live = true
	board[4][1].live = true

	*/
	board[1][6].live = true
	board[2][6].live = true
	board[3][6].live = true
	board[4][6].live = true

	for i := 0; i < 100; i++ {
		board.lookAround()
		board.print()
		board.checkLiving()
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}
