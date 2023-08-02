package draw

type (
	Color string
	Image string
	Type  int
)

const (
	Default Color = "\033[0m"
	White   Color = "\033[97m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
)

const (
	Square         Image = "[]"
	HorizontalLine Image = "=="
	VerticalLine   Image = "||"
	Space          Image = "  "
)

const (
	Edge  Type = 0
	Empty Type = 1
)
