package gform

import (
	"w32"
)

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

type LVEndLabelEditEventArg struct {
	Item *w32.LVITEM
}