package gform

import (
	"w32"
	"w32/user32"
)

type CustomControl struct {
	ControlBase
}

func (this *CustomControl) Init(parent Controller) {
	RegClassOnlyOnce("gform_customcontrol")
	this.hwnd = CreateWindow("gform_customcontrol", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE)
    this.ControlBase.init(parent)
	RegMsgHandler(this.hwnd, this)
}

func (this *CustomControl) WndProc(hwnd w32.HWND, msg uint, wparam, lparam uintptr) uintptr {
    return user32.DefWindowProc(hwnd, msg, wparam, lparam)
}
