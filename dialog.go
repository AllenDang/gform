package gform

import (
    "github.com/AllenDang/w32"
)

type Dialog struct {
    Form

    isModal  bool
    template *uint16

    Data interface{}

    onLoad         EventManager
    onOK, onCancel EventManager
}

func NewDialogFromResId(parent Controller, resId uint) *Dialog {
    d := new(Dialog)

    d.isForm = true
    d.isModal = false
    d.template = w32.MakeIntResource(uint16(resId))

    d.OnOK().Bind(dlg_OnOK)
    d.OnCancel().Bind(dlg_OnCancel)

    if parent != nil {
        d.parent = parent
    }

    return d
}

// internal event handlers
func dlg_OnOK(arg *EventArg) {
    if d, ok := arg.Sender().(*Dialog); ok {
        d.Close(w32.IDOK)
    }
}

func dlg_OnCancel(arg *EventArg) {
    if d, ok := arg.Sender().(*Dialog); ok {
        d.Close(w32.IDCANCEL)
    }
}

// Events
func (this *Dialog) OnLoad() *EventManager {
    return &this.onLoad
}

func (this *Dialog) OnOK() *EventManager {
    return &this.onOK
}

func (this *Dialog) OnCancel() *EventManager {
    return &this.onCancel
}

// Public methods
func (this *Dialog) Show() {
    this.ShowWithData(nil)
}

func (this *Dialog) ShowModal() int {
    return this.ShowModalWithData(nil)
}

func (this *Dialog) ShowWithData(data interface{}) {
    var parentHwnd w32.HWND
    if this.Parent() != nil {
        parentHwnd = this.Parent().Handle()
    }

    if this.Handle() == 0 {
        gDialogWaiting = this
        this.hwnd = w32.CreateDialog(GetAppInstance(), this.template, parentHwnd, GeneralWndprocCallBack)
        this.Data = data
        if ico, err := NewIconFromResource(GetAppInstance(), 101); err == nil {
            this.SetIcon(0, ico)
        }
    }
    this.Form.Show()
}

func (this *Dialog) ShowModalWithData(data interface{}) (result int) {
    this.isModal = true
    this.Data = data

    var parentHwnd w32.HWND
    if this.Parent() != nil {
        parentHwnd = this.Parent().Handle()
    }

    gDialogWaiting = this
    if result = w32.DialogBox(GetAppInstance(), this.template, parentHwnd, GeneralWndprocCallBack); result == -1 {
        panic("Failed to create modal dialog box")
    }

    return result
}

func (this *Dialog) Close(result int) {
    this.onClose.Fire(NewEventArg(this, nil))

    if this.isModal {
        w32.EndDialog(this.hwnd, uintptr(result))
    } else {
        w32.DestroyWindow(this.hwnd)
    }

    UnRegMsgHandler(this.hwnd)
}

func (this *Dialog) PreTranslateMessage(msg *w32.MSG) bool {
    if msg.Message >= w32.WM_KEYFIRST && msg.Message <= w32.WM_KEYLAST {
        if !this.isModal && w32.IsDialogMessage(this.hwnd, msg) {
            return true
        }
    }

    return false
}

func (this *Dialog) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    switch msg {
    case w32.WM_INITDIALOG:
        gDialogWaiting = nil
        this.onLoad.Fire(NewEventArg(this, nil))
    case w32.WM_COMMAND:
        switch w32.LOWORD(uint(wparam)) {
        case w32.IDOK:
            this.onOK.Fire(NewEventArg(this, nil))
            return w32.TRUE
        case w32.IDCANCEL:
            this.onCancel.Fire(NewEventArg(this, nil))
            return w32.TRUE
        }
    case w32.WM_CLOSE:
        this.Close(w32.IDCANCEL)
    case w32.WM_DESTROY:
        if this.parent == nil {
            Exit()
        }
    }
    return w32.FALSE
}
