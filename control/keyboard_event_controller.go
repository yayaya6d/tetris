package control

import (
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
			if controlEvent, exist := controlEventMap[event.Key]; exist {
				pollingChan <- controlEvent
			}
		}
		termbox.Flush()
	}
}
