package main

import (
	"fmt"

	"github.com/yayaya6d/tetris/console"
)

func main() {
	fmt.Println("hello tetris")

	// var c = make(chan int)

	// km := control.NewKeyBoardEventController()
	// km.StartPollingControlEvent(c)

	// <-c

	// cells := [][]draw.Cell{
	// 	{
	// 		{
	// 			Color:   draw.Blue,
	// 			Content: draw.HorizontalLine,
	// 		},
	// 		{
	// 			Color:   draw.Yellow,
	// 			Content: draw.VerticalLine,
	// 		},
	// 	},
	// 	{
	// 		{
	// 			Color:   draw.Red,
	// 			Content: draw.HorizontalLine,
	// 		},
	// 		{
	// 			Color:   draw.Green,
	// 			Content: draw.VerticalLine,
	// 		},
	// 	},
	// }

	// dr := draw.NewTetrisGameDrawer()
	// dr.Draw(cells)

	g := console.NewTetrisGameConsole()
	g.StartGame()
}
