package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameLogic(t *testing.T) {
	game, err := NewGame(5, 8)
	assert.NoError(t, err)

	game.board[1][6].live = true
	game.board[2][6].live = true
	game.board[3][6].live = true
	game.board[4][6].live = true

	game.lookAround()
	game.print()
	game.checkLiving()
	fmt.Println()

	game.lookAround()
	game.print()

	assert.Equal(t, true, game.board[2][5].live)
	assert.Equal(t, true, game.board[2][6].live)
	assert.Equal(t, true, game.board[2][7].live)

	assert.Equal(t, true, game.board[3][5].live)
	assert.Equal(t, true, game.board[3][6].live)
	assert.Equal(t, true, game.board[3][7].live)
}

func TestNewGame(t *testing.T) {
	game, err := NewGame(5, 8)
	assert.NotNil(t, game)
	assert.NoError(t, err)

	game, err = NewGame(0, 0)
	assert.Nil(t, game)
	assert.Error(t, err)

}
