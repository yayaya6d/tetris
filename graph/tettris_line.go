package graph

type line struct {
	tetrisGraph
}

func getCurrentTetrisLineGraph(i int) [][]int {
	graphs := [][][]int{
		{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
		},
		{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
		},
		{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
		},
	}

	if i >= len(graphs) {
		panic("there is a bug!! out of bound!!")
	}

	return graphs[i]
}

func (l *line) GetCurGraph() [][]int {
	return getCurrentTetrisLineGraph(l.curGraphStatus)
}
