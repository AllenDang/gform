package gform

import (
    "unsafe"
    "w32"
    "w32/user32"
)

type Form struct {
    ControlBase

    isDialog bool
}

func NewForm(parent Controller) *Form {
    f := new(Form)
    f.init(parent)

    f.SetFont(DefaultFont)
    f.SetCaption("Form")

    return f
}

func (this *Form) init(parent Controller) {
    RegClassOnlyOnce("gform_Form")

    this.isForm = true
    this.isDialog = false
    this.hwnd = CreateWindow("gform_Form", parent, w32.WS_EX_CLIENTEDGE, w32.WS_OVERLAPPEDWINDOW)
    this.ControlBase.init(parent)

    RegMsgHandler(this)
}

// Public methods
func (this *Form) Center() {
    sWidth := user32.GetSystemMetrics(w32.SM_CXFULLSCREEN)
    sHeight := user32.GetSystemMetrics(w32.SM_CYFULLSCREEN)

    if sWidth != 0 && sHeight != 0 {
        w, h := this.Size()
        this.SetPos((sWidth/2)-(w/2), (sHeight/2)-(h/2))
    }
}

// IconType: 1 - ICON_BIG; 0 - ICON_SMALL
func (this *Form) SetIcon(iconType int, icon *Icon) {
    if iconType > 1 {
        panic("IconType is invalid")
    }

    user32.SendMessage(this.hwnd, w32.WM_SETICON, uintptr(iconType), uintptr(icon.Handle()))
}

func (this *Form) SetMaxButtonEnabled(b bool) {
    ToggleStyle(this.hwnd, b, w32.WS_MAXIMIZEBOX)
}

func (this *Form) SetMinButtonEnabled(b bool) {
    ToggleStyle(this.hwnd, b, w32.WS_MINIMIZEBOX)
}

func (this *Form) SetSizable(b bool) {
    ToggleStyle(this.hwnd, b, w32.WS_THICKFRAME)
}

func (this *Form) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    switch msg {
    case w32.WM_NOTIFY: //Reflect
        nm := (*w32.NMHDR)(unsafe.Pointer(lparam))
        if msgHandler := GetMsgHandler(nm.HwndFrom); msgHandler != nil {
            return msgHandler.WndProc(msg, wparam, lparam)
        }
    case w32.WM_COMMAND: //Reflect
        if lparam != 0 { //Control
            h := w32.HWND(lparam)
            if msgHandler := GetMsgHandler(h); msgHandler != nil {
                return msgHandler.WndProc(msg, wparam, lparam)
            }
        }
    case w32.WM_CLOSE:
        user32.DestroyWindow(this.hwnd)
    case w32.WM_DESTROY:
        user32.PostQuitMessage(0)
    }

    return user32.DefWindowProc(this.hwnd, msg, wparam, lparam)
}
