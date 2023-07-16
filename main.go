package main

import (
	"fmt"

	"github.com/yayaya6d/tetris/control"
)

func main() {
	fmt.Print("hello tetris")

	var c = make(chan int)

	km := control.NewKeyBoardEventController()
	km.StartPollingControlEvent(c)

	<-c
}
