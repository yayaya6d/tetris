package graph

import "math/rand"

var graphMap = []func() TetrisGraph{
	func() TetrisGraph {
		return &line{}
	},
}

func generateRandomGraph() TetrisGraph {
	n := rand.Int() % len(graphMap)
	ret := graphMap[n]()
	return ret
}
