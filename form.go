package gform

import (
    "github.com/AllenDang/w32"
)

type Form struct {
    ControlBase

    isDialog   bool
    isDragMove bool
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
    this.isDragMove = false
    this.hwnd = CreateWindow("gform_Form", parent, w32.WS_EX_CLIENTEDGE, w32.WS_OVERLAPPEDWINDOW)
    this.ControlBase.init(parent)

    RegMsgHandler(this)
}

// Public methods
func (this *Form) Center() {
    sWidth := w32.GetSystemMetrics(w32.SM_CXFULLSCREEN)
    sHeight := w32.GetSystemMetrics(w32.SM_CYFULLSCREEN)

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

    w32.SendMessage(this.hwnd, w32.WM_SETICON, uintptr(iconType), uintptr(icon.Handle()))
}

func (this *Form) EnableMaxButton(b bool) {
    ToggleStyle(this.hwnd, b, w32.WS_MAXIMIZEBOX)
}

func (this *Form) EnableMinButton(b bool) {
    ToggleStyle(this.hwnd, b, w32.WS_MINIMIZEBOX)
}

func (this *Form) EnableSizable(b bool) {
    ToggleStyle(this.hwnd, b, w32.WS_THICKFRAME)
}

func (this *Form) EnableDragMove(b bool) {
    this.isDragMove = b
}

func (this *Form) EnableTopMost(b bool) {
    tag := w32.HWND_NOTOPMOST
    if b {
        tag = w32.HWND_TOPMOST
    }

    w32.SetWindowPos(this.hwnd, tag, 0, 0, 0, 0, w32.SWP_NOMOVE|w32.SWP_NOSIZE)
}

func (this *Form) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    switch msg {
    case w32.WM_LBUTTONDOWN:
        if this.isDragMove {
            w32.ReleaseCapture()
            w32.SendMessage(this.hwnd, w32.WM_NCLBUTTONDOWN, w32.HTCAPTION, 0)
        }
    case w32.WM_CLOSE:
        w32.DestroyWindow(this.hwnd)
    case w32.WM_DESTROY:
        w32.PostQuitMessage(0)
    }

    return w32.DefWindowProc(this.hwnd, msg, wparam, lparam)
}
