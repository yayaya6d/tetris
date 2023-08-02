package console

import (
	"fmt"
	"time"

	"github.com/yayaya6d/tetris/control"
	"github.com/yayaya6d/tetris/draw"
	"github.com/yayaya6d/tetris/processor"
)

const (
	frameTime    = 33 * time.Millisecond
	refreshCount = 30
)

type Console interface {
	ShowHelp()
	StartGame()
}

type tetrisGameConsole struct {
	drawer          draw.Drawer
	eventController control.EventController
	gameProcessor   *processor.TetrisGameProcessor
}

func NewTetrisGameConsole() Console {
	return &tetrisGameConsole{
		drawer:          draw.NewTetrisGameDrawer(),
		eventController: control.NewKeyBoardEventController(),
		gameProcessor:   &processor.TetrisGameProcessor{},
	}
}

func (c *tetrisGameConsole) ShowHelp() {

}

func (c *tetrisGameConsole) StartGame() {
	c.gameProcessor.Init()

	running := true
	eventReceivedChan := make(chan int)
	c.eventController.StartPollingControlEvent(eventReceivedChan)

	refreshCounter := 0

	for running {
		time.Sleep(frameTime)
		refreshCounter++

		select {
		case evt := <-eventReceivedChan:
			if evt == control.Exit {
				running = false
				fmt.Println("exit console")
			} else {
				c.processEvent(evt)
			}
		default:
			if refreshCounter == refreshCount {
				c.drawer.Draw(c.gameProcessor.GetCurrentCells())
				refreshCounter = 0
			}
		}
	}
}

func (c *tetrisGameConsole) processEvent(evt int) {
	fmt.Printf("receive %d\n", evt)
}
