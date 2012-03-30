package gform

import (
	"github.com/AllenDang/w32"
	"github.com/AllenDang/w32/user32"
)

type CustomControl struct {
	W32Control

    ClassName      string
    ExStyle, Style uint
}

func (this *CustomControl) Init(parent Controller) {
    if len(this.ClassName) == 0 {
        panic("Class name of custom control cannot be nil")
    }

	RegClassOnlyOnce(this.ClassName)

    if this.Style == 0 {
        this.Style = w32.WS_CHILD | w32.WS_VISIBLE
    }
	this.hwnd = CreateWindow(this.ClassName, parent, this.ExStyle, this.Style)
    this.isMouseLeft = true
    this.ControlBase.init(parent)
}

func (this *CustomControl) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    sender := GetMsgHandler(this.hwnd)
    switch msg {
    case w32.WM_CREATE:
        internalTrackMouseEvent(this.hwnd)
        this.onCreate.Fire(NewEventArg(sender, nil))
    case w32.WM_CLOSE:
        this.onClose.Fire(NewEventArg(sender, nil))
    case w32.WM_MOUSEMOVE:
        if this.isMouseLeft {
            this.onMouseHover.Fire(NewEventArg(sender, nil))
            internalTrackMouseEvent(this.hwnd)
            this.isMouseLeft = false
        }
    case w32.WM_MOUSELEAVE:
        this.onMouseLeave.Fire(NewEventArg(sender, nil))
        this.isMouseLeft = true
    }
    return user32.DefWindowProc(this.hwnd, msg, wparam, lparam)
}
