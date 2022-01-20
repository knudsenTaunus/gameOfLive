package main

type Node struct {
	width    int
	height   int
	live     bool
	willLive bool
	counter  int
}

// Each node looks what is the status of all its neighbours ...
func (n *Node) lookAround(board *Game) {
	count := 0

	south := board.getNode(n.height+1, n.width)
	if south != nil && south.live {
		count++
	}
	southeast := board.getNode(n.height+1, n.width+1)
	if southeast != nil && southeast.live {
		count++
	}
	southwest := board.getNode(n.height+1, n.width-1)
	if southwest != nil && southwest.live {
		count++
	}
	east := board.getNode(n.height, n.width+1)
	if east != nil && east.live {
		count++
	}
	west := board.getNode(n.height, n.width-1)
	if west != nil && west.live {
		count++
	}
	north := board.getNode(n.height-1, n.width)
	if north != nil && north.live {
		count++
	}
	northwest := board.getNode(n.height-1, n.width-1)
	if northwest != nil && northwest.live {
		count++
	}
	northeast := board.getNode(n.height-1, n.width+1)
	if northeast != nil && northeast.live {
		count++
	}

	// ... here the rules of the game are applied
	if n.live {
		if count < 2 {
			n.willLive = false
			return
		}

		if count > 3 {
			n.willLive = false
			return
		}
		n.willLive = true
		return
	}
	if count == 3 {
		n.willLive = true
		return
	}
}
