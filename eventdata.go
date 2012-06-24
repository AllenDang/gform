package gform

import (
	"github.com/AllenDang/w32"
)

type RawMsg struct {
    Hwnd           w32.HWND
    Msg            uint
    WParam, LParam uintptr
}

type MouseEventData struct {
    X, Y   int
    Button int
    Wheel  int
}

type DropFilesEventData struct {
    X, Y  int
	Files []string
}

type PaintEventData struct {
	Canvas *Canvas
}

type LVEndLabelEditEventData struct {
	Item *w32.LVITEM
}

type LVDBLClickEventData struct {
	NmItem *w32.NMITEMACTIVATE
}

type KeyUpEventData struct {
	VKey, Code int
}

type SizeEventData struct {
	Type uint
	X, Y int
}
