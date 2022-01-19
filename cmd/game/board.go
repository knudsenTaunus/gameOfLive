package main

type Board [HEIGHT][WIDTH]*Node

func NewBoard() *Board {
	b := &Board{}
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			b[i][j] = &Node{
				width:    j + 1,
				height:   i + 1,
				willLive: false,
			}
		}
	}
	return b
}

func (b *Board) checkLiving() {
	for _, h := range b {
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

func (b *Board) lookAround() {
	for _, h := range b {
		for _, node := range h {
			node.lookAround(b)
		}
	}
}

func (b *Board) getNode(i, j int) *Node {
	if i-1 >= HEIGHT || i-1 < 0 {
		return nil
	}

	if j-1 >= WIDTH || j-1 < 0 {
		return nil
	}

	return b[i-1][j-1]
}

func (b *Board) print() {
	for _, h := range b {
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
