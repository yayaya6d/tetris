package draw

type Cell struct {
	Content Image
	Color   Color
	Type    Type
}

type Drawer interface {
	Draw([][]Cell)
}
