package gform

import (
    "w32"
    "w32/user32"
)

type W32Control struct {
    ControlBase
    originalWndProc uintptr
}

func (this *W32Control) init(className string, parent Controller, exstyle, style uint) {
    this.hwnd = CreateWindow(className, parent, exstyle, style)
    if this.hwnd == 0 {
        panic("cannot create window for " + className)
    }
    this.originalWndProc = user32.SetWindowLongPtr(this.hwnd, w32.GWLP_WNDPROC, GeneralWndprocCallBack)
    this.ControlBase.init(parent)
}

func (this *W32Control) attach(parent Controller, dlgItemID int) {
    if parent == nil {
        panic("parent cannot be nil")
    }

    if this.hwnd = user32.GetDlgItem(parent.Handle(), dlgItemID); this.hwnd == 0 {
        panic("hwnd cannot be nil")
    }

    this.originalWndProc = user32.SetWindowLongPtr(this.hwnd, w32.GWLP_WNDPROC, GeneralWndprocCallBack)
    this.ControlBase.init(parent)
}

func (this *W32Control) WndProc(hwnd w32.HWND, msg uint, wparam, lparam uintptr) uintptr {
    return user32.CallWindowProc(this.originalWndProc, this.hwnd, msg, wparam, lparam)
}
