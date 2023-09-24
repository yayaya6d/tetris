package graph

type GraphFactory interface {
	CreateNewGraph() TetrisGraph
}

type graphFactory struct {
	queueSize int
	queue     []TetrisGraph
}

func NewGraphFactory(queueSize int) GraphFactory {
	ret := &graphFactory{
		queueSize: queueSize,
		queue:     make([]TetrisGraph, queueSize),
	}
	ret.initQueue()

	return ret
}

func (f *graphFactory) initQueue() {
	for i := 0; i < f.queueSize; i++ {
		f.queue[i] = generateRandomGraph()
	}
}

func (f *graphFactory) CreateNewGraph() TetrisGraph {
	ret := f.queue[0]

	f.queue = f.queue[1:len(f.queue)]
	f.queue[f.queueSize-1] = generateRandomGraph()

	return ret
}
