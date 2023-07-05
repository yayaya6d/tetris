package draw

import "fmt"

type tetrixGameDrawer struct {
}

func NewTetrisGameDrawer() Drawer {
	return &tetrixGameDrawer{}
}

func (d *tetrixGameDrawer) Draw(cells [][]Cell) {
	for _, cs := range cells {
		for _, c := range cs {
			fmt.Printf("%s%s", c.Color, c.Content)
		}
		fmt.Print("\n")
	}
}
