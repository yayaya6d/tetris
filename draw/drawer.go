package draw

type Cell struct {
	Content Image
	Color   Color
}

type Drawer interface {
	Draw([][]Cell)
}
