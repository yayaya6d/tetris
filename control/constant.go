package control

import "github.com/nsf/termbox-go"

const (
	Spin int = iota
	Down
	Left
	Right
	Exit
)

var controlEventMap = map[termbox.Key]int{
	termbox.KeyArrowUp:    Spin,
	termbox.KeyArrowDown:  Down,
	termbox.KeyArrowLeft:  Left,
	termbox.KeyArrowRight: Right,
	termbox.KeyEsc:        Exit,
}
