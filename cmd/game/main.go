package main

import (
	"fmt"
	"time"
)

const (
	width  = 8
	height = 5
)

type Board [height][width]*Node

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) updateLive() {
	for _, h := range b {
		for _, node := range h {
			if node.counter < 2 {
				node.willLive = false
				node.sign = "."
			}

			if node.counter > 3 {
				node.willLive = false
				node.sign = "."
			}

			if node.counter == 2 || node.counter == 3 {
				if node.live {
					node.willLive = true
					node.sign = "*"
				}
				if node.live == false && node.counter == 3 {
					node.willLive = true
					node.sign = "*"
				}
			}
		}
	}
}

func (b *Board) resetLive() {
	for _, h := range b {
		for _, node := range h {
			if node.willLive {
				node.live = true
				node.willLive = false
				node.sign = "*"
				node.counter = 0
				continue
			}
			if node.live {
				node.live = false
				node.willLive = false
				node.sign = "."
				node.counter = 0
				continue
			}
			node.live = false
			node.willLive = false
			node.counter = 0
		}
	}
}

func (b *Board) getNode(i, j int) *Node {
	if i-1 >= height || i-1 < 0 {
		return nil
	}

	if j-1 >= width || j-1 < 0 {
		return nil
	}

	return b[i-1][j-1]
}

func (b *Board) print() {
	for _, h := range b {
		for _, node := range h {
			print(node.sign, "\t")
		}
		println()
	}

}

func main() {
	board := NewBoard()
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			board[i][j] = &Node{
				width:      j + 1,
				height:     i + 1,
				sign:       ".",
				willLive:   false,
				neighbours: [8]*Node{},
			}
		}
	}

	board[0][1].live = true
	board[0][1].sign = "*"
	board[0][5].live = true
	board[0][5].sign = "*"
	board[1][3].live = true
	board[1][3].sign = "*"
	board[2][1].live = true
	board[2][1].sign = "*"
	board[3][3].live = true
	board[3][3].sign = "*"
	board[4][1].live = true
	board[4][1].sign = "*"

	board.print()
	fmt.Println("---")

	for i := 0; i < 100; i++ {
		fertilize(board)
		board.updateLive()
		board.print()
		board.resetLive()
		fmt.Println("---")
		time.Sleep(1 * time.Second)
	}
}

func fertilize(board *Board) {
	for _, h := range board {
		for _, node := range h {
			if node.live {
				node.sprayAround(board)
			}
		}
	}
}
