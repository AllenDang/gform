package gform

import (
	"w32"
	"w32/user32"
)

type CustomControl struct {
	W32Control
}

func (this *CustomControl) Init(parent Controller) {
	RegClassOnlyOnce("gform_customcontrol")
	this.hwnd = CreateWindow("gform_customcontrol", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE)
    this.isMouseLeft = true
    this.ControlBase.init(parent)
}

func (this *CustomControl) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    sender := GetMsgHandler(this.hwnd)
    switch msg {
    case w32.WM_CREATE:
        internalTrackMouseEvent(this.hwnd)
        this.onCreate.Fire(sender)
    case w32.WM_MOUSEMOVE:
        if this.isMouseLeft {
            this.onMouseHover.Fire(sender)
            internalTrackMouseEvent(this.hwnd)
            this.isMouseLeft = false
        }
    case w32.WM_MOUSELEAVE:
        this.onMouseLeave.Fire(sender)
        this.isMouseLeft = true
    }
    return user32.DefWindowProc(this.hwnd, msg, wparam, lparam)
}
