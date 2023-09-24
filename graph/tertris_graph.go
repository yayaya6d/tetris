package graph

const tetrisGraphNumCount = 4

type TetrisGraph interface {
	Init(x, y int)
	Spin()
	SpinBack()
	MoveUp()
	MoveDown()
	MoveLeft()
	MoveRight()
	GetCurPos() Positition
	GetCurGraph() [][]int
}

type tetrisGraph struct {
	curPos         Positition
	curGraphStatus int
}

func (g *tetrisGraph) Init(x, y int) {
	g.curPos.x = x
	g.curPos.y = y
	g.curGraphStatus = 0
}

func (g *tetrisGraph) Spin() {
	g.curGraphStatus++
	if g.curGraphStatus >= tetrisGraphNumCount {
		g.curGraphStatus = 0
	}
}

func (g *tetrisGraph) SpinBack() {
	g.curGraphStatus--
	if g.curGraphStatus < 0 {
		g.curGraphStatus = tetrisGraphNumCount
	}
}

func (g *tetrisGraph) MoveUp() {

}

func (g *tetrisGraph) MoveDown() {
	g.curPos.y--
}

func (g *tetrisGraph) MoveLeft() {
	g.curPos.x--
}

func (g *tetrisGraph) MoveRight() {
	g.curPos.y++
}

func (g *tetrisGraph) GetCurPos() Positition {
	return g.curPos
}
