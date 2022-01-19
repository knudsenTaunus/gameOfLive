package main

type Node struct {
	width      int
	height     int
	sign       string
	live       bool
	willLive   bool
	counter    int
	neighbours [8]*Node
}

func (n *Node) sprayAround(board *Board) {
	south := board.getNode(n.height+1, n.width)
	if south != nil {
		south.counter++
	}
	southeast := board.getNode(n.height+1, n.width+1)
	if southeast != nil {
		southeast.counter++
	}
	southwest := board.getNode(n.height+1, n.width-1)
	if southwest != nil {
		southwest.counter++
	}
	east := board.getNode(n.height, n.width+1)
	if east != nil {
		east.counter++
	}
	west := board.getNode(n.height, n.width-1)
	if west != nil {
		west.counter++
	}
	north := board.getNode(n.height-1, n.width)
	if north != nil {
		north.counter++
	}
	northwest := board.getNode(n.height-1, n.width-1)
	if northwest != nil {
		northwest.counter++
	}
	northeast := board.getNode(n.height-1, n.width+1)
	if northeast != nil {
		northeast.counter++
	}
}
