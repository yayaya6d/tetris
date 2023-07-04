package control

import (
	"fmt"
	"sync"

	"github.com/nsf/termbox-go"
)

var singletonKeyBoardEventController *keyboardEventController
var once sync.Once

func NewKeyBoardEventController() EventController {
	once.Do(func() {
		singletonKeyBoardEventController = &keyboardEventController{
			startPolling: false,
		}
	})
	return singletonKeyBoardEventController
}

type keyboardEventController struct {
	startPolling bool
}

func (c *keyboardEventController) StartPollingControlEvent(pollingChan chan<- int) {
	if !c.startPolling {
		c.startPolling = true
		go c.pollingKeyboardEvent(pollingChan)
	}
}

func (c *keyboardEventController) pollingKeyboardEvent(pollingChan chan<- int) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	running := true

	for running {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyArrowDown {
				fmt.Println("Press Down")
			} else if event.Key == termbox.KeyArrowUp {
				fmt.Println("Press Up")
			} else if event.Key == termbox.KeyArrowLeft {
				fmt.Println("Press Left")
			} else if event.Key == termbox.KeyArrowRight {
				fmt.Println("Press Right")
			} else if event.Key == termbox.KeyEsc {
				running = false
				pollingChan <- 1
				fmt.Println("Exit")
			} else {
				fmt.Println("Dont care")
			}
			termbox.Flush()
		default:
			fmt.Println("Dont care")
			termbox.Flush()
		}
	}
}
