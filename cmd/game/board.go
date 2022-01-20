package main

import "errors"

type Game struct {
	height int
	width  int
	board  [][]*Node
}

// NewGame Constructor receives height and width and returns a new Game with a board
func NewGame(height, width int) (*Game, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("dimensions must be greater then 0")
	}

	board := make([][]*Node, height, width)
	for i := 0; i < height; i++ {
		board[i] = make([]*Node, width)
		for j := 0; j < width; j++ {
			board[i][j] = &Node{
				width:    j + 1,
				height:   i + 1,
				willLive: false,
			}
		}
	}

	game := &Game{
		height: height,
		width:  width,
		board:  board,
	}

	return game, nil
}

// This function updates the cells according to their status
func (b *Game) checkLiving() {
	for _, h := range b.board {
		for _, node := range h {
			if node.willLive {
				node.live = true
				continue
			}
			node.live = false
			continue
		}
	}
}

func (b *Game) lookAround() {
	for _, h := range b.board {
		for _, node := range h {
			node.lookAround(b)
		}
	}
}

// A helper function which safely gets a node or nil if the boundaries are reached
func (b *Game) getNode(i, j int) *Node {
	if i-1 >= b.height || i-1 < 0 {
		return nil
	}

	if j-1 >= b.width || j-1 < 0 {
		return nil
	}

	return b.board[i-1][j-1]
}

// Prints each cell of the board according to its status
func (b *Game) print() {
	for _, h := range b.board {
		for _, node := range h {
			if node.live {
				print("*", "\t")
				continue
			}
			print(".", "\t")
		}
		println()
	}

}
