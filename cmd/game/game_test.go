package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameLogic(t *testing.T) {
	board := NewBoard()

	board[1][6].live = true
	board[2][6].live = true
	board[3][6].live = true
	board[4][6].live = true

	board.lookAround()
	board.print()
	board.checkLiving()
	fmt.Println()

	board.lookAround()
	board.print()

	assert.Equal(t, true, board[2][5].live)
	assert.Equal(t, true, board[2][6].live)
	assert.Equal(t, true, board[2][7].live)

	assert.Equal(t, true, board[3][5].live)
	assert.Equal(t, true, board[3][6].live)
	assert.Equal(t, true, board[3][7].live)

}
