package gform

import (
    "unsafe"
    "w32"
    "w32/user32"
)

type Dialog struct {
    Form

    isModal  bool
    result   int
    template *uint16

    onLoad GeneralEventManager
}

func NewDialogFromTemplate(parent Controller, template *uint16) *Dialog {
    d := new(Dialog)

    d.isForm = true
    d.isModal = false
    d.template = template

    if parent != nil {
        d.parent = parent
    }

    return d
}

func (this *Dialog) OnLoad() *GeneralEventManager {
    return &this.onLoad
}

func (this *Dialog) Show() {
    var parentHwnd w32.HWND
    if this.Parent() != nil {
        parentHwnd = this.Parent().Handle()
    }

    this.hwnd = user32.CreateDialog(GetAppInstance(), this.template, parentHwnd, GeneralWndprocCallBack)
    RegMsgHandler(this)

    this.onLoad.Fire(this)

    this.Form.Show()
}

func (this *Dialog) ShowModal() int {
    this.isModal = true

    var parentHwnd w32.HWND
    if this.Parent() != nil {
        parentHwnd = this.Parent().Handle()
    }

    if this.result = user32.DialogBox(GetAppInstance(), this.template, parentHwnd, GeneralWndprocCallBack); this.result == -1 {
        panic("Failed to create modal dialog box")
    }
    
    return this.result
}

func (this *Dialog) Close() {
    UnRegMsgHandler(this.hwnd)
    if this.isModal {
        user32.EndDialog(this.hwnd, uintptr(this.result))
    } else {
        user32.DestroyWindow(this.hwnd)
    }
}

func (this *Dialog) PreTranslateMessage(msg *w32.MSG) bool {
    if msg.Message >= w32.WM_KEYFIRST && msg.Message <= w32.WM_KEYLAST {
        if !this.isModal && user32.IsDialogMessage(this.hwnd, msg) {
            return true
        }
    }

    return false
}

func (this *Dialog) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    switch msg {
    case w32.WM_NOTIFY:
        nm := (*w32.NMHDR)(unsafe.Pointer(lparam))
        if msgHandler := GetMsgHandler(nm.HwndFrom); msgHandler != nil {
            ret := msgHandler.WndProc(msg, wparam, lparam)
            if ret != 0 {
                user32.SetWindowLong(this.hwnd, w32.DWL_MSGRESULT, uint32(ret))
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
                    user32.SetWindowLong(this.hwnd, w32.DWL_MSGRESULT, uint32(ret))
                    return w32.TRUE
                }
            }
        }
        switch w32.LOWORD(uint(wparam)) {
        case w32.IDCANCEL:
            user32.DestroyWindow(this.hwnd)
            return w32.TRUE
        }
    }
    return w32.FALSE
}
