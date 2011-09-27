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

func NewFormFromTemplate(parent Controller, template *uint16) *Form {
    f := new(Form)
    f.initFromTemplate(parent, template)

    return f
}

func (this *Form) init(parent Controller) {
    RegClassOnlyOnce("gform_Form")

    this.isForm = true
    this.isDialog = false
    this.hwnd = CreateWindow("gform_Form", parent, w32.WS_EX_CLIENTEDGE, w32.WS_OVERLAPPEDWINDOW)
    this.ControlBase.init(parent)

    RegMsgHandler(this.hwnd, this)
}

func (this *Form) initFromTemplate(parent Controller, template *uint16) {
    this.isForm = true
    this.isDialog = true
    var parentHwnd w32.HWND = 0
    if parent != nil {
        parentHwnd = parent.Handle()
    }
    this.hwnd = user32.CreateDialog(GetAppInstance(), template, parentHwnd, GeneralWndprocCallBack)
    this.ControlBase.init(parent)

    RegMsgHandler(this.hwnd, this)
}

func (this *Form) Center() {
    sWidth := user32.GetSystemMetrics(w32.SM_CXFULLSCREEN)
    sHeight := user32.GetSystemMetrics(w32.SM_CYFULLSCREEN)

    if sWidth != 0 && sHeight != 0 {
        w, h := this.Size()
        this.SetPos((sWidth/2)-(w/2), (sHeight/2)-(h/2))
    }
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
    if this.isDialog {
        switch msg {
        case w32.WM_INITDIALOG:
            return w32.TRUE
        case w32.WM_NOTIFY:
            nm := (*w32.NMHDR)(unsafe.Pointer(lparam))
            if msgHandler := GetMsgHandler(nm.HwndFrom); msgHandler != nil {
                ret := msgHandler.WndProc(msg, wparam, lparam)
                if ret != 0 {
                    user32.SetWindowLong(this.Handle(), w32.DWL_MSGRESULT, uint32(ret))
                    return w32.TRUE
                }
            }
        case w32.WM_DESTROY:
            user32.PostQuitMessage(0)
        case w32.WM_COMMAND:
            if lparam != 0 { //Control
                h := w32.HWND(lparam)
                if msgHandler := GetMsgHandler(h); msgHandler != nil {
                    ret := msgHandler.WndProc(msg, wparam, lparam)
                    if ret != 0 {
                        user32.SetWindowLong(this.Handle(), w32.DWL_MSGRESULT, uint32(ret))
                        return w32.TRUE
                    }
                }
            }
            switch w32.LOWORD(uint(wparam)) {
            case w32.IDCANCEL:
                user32.DestroyWindow(this.Handle())
                return w32.TRUE
            }
        }
        return w32.FALSE
    } else {
        switch msg {
        case w32.WM_NOTIFY:
            nm := (*w32.NMHDR)(unsafe.Pointer(lparam))
            if msgHandler := GetMsgHandler(nm.HwndFrom); msgHandler != nil {
                return msgHandler.WndProc(msg, wparam, lparam)
            }
        case w32.WM_COMMAND:
            if lparam != 0 { //Control
                h := w32.HWND(lparam)
                if msgHandler := GetMsgHandler(h); msgHandler != nil {
                    return msgHandler.WndProc(msg, wparam, lparam)
                }
            }
        case w32.WM_CLOSE:
            user32.DestroyWindow(this.Handle())
        case w32.WM_DESTROY:
            user32.PostQuitMessage(0)
        }

        return user32.DefWindowProc(this.Handle(), msg, wparam, lparam)
    }

    return 0
}
