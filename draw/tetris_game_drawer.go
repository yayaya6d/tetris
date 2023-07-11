package draw

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type tetrixGameDrawer struct {
}

func NewTetrisGameDrawer() Drawer {
	return &tetrixGameDrawer{}
}

func (d *tetrixGameDrawer) Draw(cells [][]Cell) {
	d.clear()
	for _, cs := range cells {
		for _, c := range cs {
			fmt.Printf("%s%s", c.Color, c.Content)
		}
		fmt.Print("\n")
	}
}

func (d *tetrixGameDrawer) clear() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
