package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// first we get the size of the board from the command line ...
	if len(os.Args[1:]) < 2 {
		log.Println("Please provide dimensions")
		return
	}

	dimensions := os.Args[1:]
	h := dimensions[0]
	w := dimensions[1]

	height, err := strconv.Atoi(h)
	if err != nil {
		log.Println(err)
		return
	}

	width, err := strconv.Atoi(w)
	if err != nil {
		log.Println(err)
		return
	}

	// ... with the dimensions we can now build our board.
	game, err := NewGame(height, width)
	if err != nil {
		log.Println(err)
		return
	}

	// We ask the user to provide the coordinates of the living cells ...
	fmt.Printf("Please enter the height and width of your next living cell within height of %d and width of %d like so: 1 3\n", height, width)
	fmt.Println("If you dont want to add cells just press enter.")

	for {
		reader := bufio.NewReader(os.Stdin)
		coords, _ := reader.ReadString('\n')
		coords = strings.TrimSuffix(coords, "\n")
		if coords == "" {
			break
		}

		widthAndHeight := strings.Split(coords, " ")
		if len(widthAndHeight) < 2 {
			fmt.Println("Seems like you just provided one number, but we need two")
			continue
		}

		cellHeight, cellErr := strconv.Atoi(widthAndHeight[0])
		if cellErr != nil {
			log.Println(cellErr)
			return
		}

		if cellHeight > height {
			log.Println("The height seems to big, please try it again")
			continue
		}

		cellWidth, cellErr := strconv.Atoi(widthAndHeight[1])
		if cellErr != nil {
			log.Println(cellErr)
			return
		}

		if cellWidth > width {
			log.Println("The width seems to big, please try it again")
			continue
		}

		game.board[cellHeight][cellWidth].live = true
		fmt.Println("Another one? Otherwise just press ENTER")
	}

	// ... when this is done, the game starts.
	for i := 0; i < 100; i++ {
		// First all the cells look around them and check the amount of living neighbours...
		game.lookAround()
		// ... the board is printed with the current status ...
		game.print()
		// ... and then the cells properties are updated.
		game.checkLiving()
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}
