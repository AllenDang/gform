package gform

import (
    "github.com/AllenDang/w32"
)

type Edit struct {
    W32Control

    onChange EventManager
}

func NewEdit(parent Controller) *Edit {
    edt := new(Edit)
    edt.init(parent)

    edt.SetFont(DefaultFont)
    edt.SetSize(200, 20)

    return edt
}

func AttachEdit(parent Controller, id int) *Edit {
    edt := new(Edit)
    edt.attach(parent, id)
    RegMsgHandler(edt)

    return edt
}

func (this *Edit) init(parent Controller) {
    this.W32Control.init("EDIT", parent, w32.WS_EX_CLIENTEDGE, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_TABSTOP|w32.ES_LEFT|w32.ES_MULTILINE)
    RegMsgHandler(this)
}

//Events
func (this *Edit) OnChange() *EventManager {
    return &this.onChange
}

//Public methods
func (this *Edit) SetReadOnly(isReadOnly bool) {
    w32.SendMessage(this.hwnd, w32.EM_SETREADONLY, uintptr(w32.BoolToBOOL(isReadOnly)), 0)
}

func (this *Edit) AddLine(text string) {
    if len(this.Caption()) == 0 {
        this.SetCaption(text)
    } else {
        this.SetCaption(this.Caption() + "\r\n" + text)
    }
}

func (this *Edit) WndProc(msg uint, wparam, lparam uintptr) uintptr {
    switch msg {
    case w32.WM_COMMAND:
        switch w32.HIWORD(uint(wparam)) {
        case w32.EN_CHANGE:
            this.onChange.Fire(NewEventArg(this, nil))
        }
    }

    return this.W32Control.WndProc(msg, wparam, lparam)
}
