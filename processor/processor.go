package processor

import "github.com/yayaya6d/tetris/draw"

type TetrisGameProcessor struct {
	cells [][]draw.Cell
}

func (p *TetrisGameProcessor) Init() {
	p.InitCells()
}

func (p *TetrisGameProcessor) InitCells() {
	p.cells = make([][]draw.Cell, height)
	for i := 0; i < height; i++ {
		p.cells[i] = make([]draw.Cell, width)
		p.cells[i][0] = draw.Cell{
			Content: draw.VerticalLine,
			Color:   draw.White,
			Type:    draw.Edge,
		}

		for j := 1; j < width-1; j++ {
			p.cells[i][j] = draw.Cell{
				Content: draw.Space,
				Color:   draw.White,
				Type:    draw.Empty,
			}
		}

		p.cells[i][width-1] = draw.Cell{
			Content: draw.VerticalLine,
			Color:   draw.White,
			Type:    draw.Edge,
		}
	}

	for i := 0; i < width; i++ {
		p.cells[height-1][i] = draw.Cell{
			Content: draw.HorizontalLine,
			Color:   draw.White,
			Type:    draw.Edge,
		}
	}
}

func (p *TetrisGameProcessor) ProcessEvent(event int) {

}

func (p *TetrisGameProcessor) GetCurrentCells() [][]draw.Cell {
	return p.cells
}
