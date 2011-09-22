package gform

type MouseEventArg struct {
    X, Y int
    Button int
    Wheel int
}

type DropFilesEventArg struct {
	X, Y int
	Files []string
}

type PaintEventArg struct {
	Canvas *Canvas
}